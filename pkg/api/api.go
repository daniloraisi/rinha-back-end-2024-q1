package api

import (
	"io"
	"net/http"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/pkg/db"
	"github.com/daniloraisi/rinha-back-end/pkg/rest"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

var healthy = false

type Config interface {
	GetDbConn() *sqlx.DB
	GetEnvironment() string
}

func SetHealthy(h bool) {
	healthy = h
}

func NewRouter(config Config, dbConfig *db.DbConfig, l *logger.Logger) *httprouter.Router {
	mux := httprouter.New()

	clientService := &db.ClientService{
		Db:       config.GetDbConn(),
		DbConfig: dbConfig,
		Logger:   l,
	}

	mux.GET("/clientes/:id/extrato", rest.GetExtrato(clientService, l))
	mux.POST("/clientes/:id/transacoes", rest.Transacionar(clientService, l))
	mux.POST("/reset-db", rest.ResetDB(config.GetDbConn(), l))

	mux.GET("/healthz", healthCheck())

	return mux
}

func healthCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if healthy {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Not Healthy")
		}
	}
}
