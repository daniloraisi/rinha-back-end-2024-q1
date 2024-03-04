package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/pkg/api"
	"github.com/daniloraisi/rinha-back-end/pkg/db"
	"github.com/jmoiron/sqlx"
)

const (
	livenessFileName = "serviceAlive"
)

var (
	healthy          = true
	LivenessFilePath = ensureLivenessFile()
)

type Config interface {
	GetDbConn() *sqlx.DB
	GetPortHTTP() uint16
	GetEnvironment() string
}

func StartAndRun(config Config, dbConfig *db.DbConfig, l *logger.Logger) {
	l.Info("iniciando servidor HTTP")

	server, portHTTP := createHTTPServer(config, api.NewRouter(config, dbConfig, l))

	chShutdownComplete := handleShutdown(server, config, l)

	listenAndServe(server, portHTTP, l)

	<-chShutdownComplete
	markDead(LivenessFilePath, l)
}

func createHTTPServer(config Config, handler http.Handler) (*http.Server, uint16) {
	portHTTP := config.GetPortHTTP()

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", portHTTP),
		Handler: handler,
	}, portHTTP
}

func handleShutdown(server *http.Server, config Config, l *logger.Logger) (chShutdownComplete chan struct{}) {
	chShutdownComplete = make(chan struct{})

	go func() { shutdown(server, config, chShutdownComplete, l) }()

	return chShutdownComplete
}

func shutdown(server *http.Server, config Config, chShutdownComplete chan struct{}, l *logger.Logger) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	l.Info("solicitação de interrupção de serviço recebida")
	api.SetHealthy(false)

	if config.GetEnvironment() == "prod" {
		time.Sleep(5 * time.Second)
	}

	healthy = false
	timeoutCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		l.Error("erro ao parar servidor HTTP", l.ErrProp(err))
	}
	l.Info("servidor HTTP desligado")

	stop()
	close(chShutdownComplete)
}

func listenAndServe(server *http.Server, portHTTP uint16, l *logger.Logger) {
	markAlive(LivenessFilePath, l)
	api.SetHealthy(true)
	l.Info("healthy check HTTP ligado", l.Prop("port", portHTTP))
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			l.Error("fatal! falha ao iniciar servidor HTTP", l.ErrProp(err))
		}
	}
}

func IsHealthy() bool {
	return healthy
}

func ensureLivenessFile() string {
	tmpPath := os.TempDir()
	if _, err := os.Stat(tmpPath); err != nil {
		if err = os.MkdirAll(tmpPath, os.ModePerm); err != nil {
			log.Fatal("fatal! diretório temporário não existe e não pode ser criado", err)
		}
	}

	return path.Join(tmpPath, livenessFileName)
}

func markAlive(filepath string, l *logger.Logger) {
	if _, err := os.Create(filepath); err != nil {
		l.Fatal("fatal! não foi possível criar o arquivo liveness", l.ErrProp(err))
	}

	l.Info("serviço marcado como executando")
}

func markDead(filepath string, l *logger.Logger) {
	if err := os.Remove(filepath); err != nil {
		l.Fatal("fatal! não foi possível excluir o arquivo liveness", l.ErrProp(err))
	}

	l.Info("serviço marcado como parado")
}
