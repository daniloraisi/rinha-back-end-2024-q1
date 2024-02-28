package main

import (
	"github.com/daniloraisi/rinha-back-end/internal/env"
	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/pkg/db"
	"github.com/daniloraisi/rinha-back-end/pkg/service"
)

func main() {
	envName := env.Panic(env.Get("ENVIRONMENT_NAME").Required().String())
	l := logger.New(envName)
	l.Info("iniciando serviço")

	l.Info("configurando serviço")

	dbConfig := getDBConfig()
	config := getServerConfig(dbConfig, l)

	defer teardown(config, l)

	service.StartAndRun(config, dbConfig, l)

	l.Info("parando serviço")
}

func teardown(config *serverConfig, l *logger.Logger) {
	l.Info("derrubando serviço")

	config.GetDbConn().Close()

	l.Info("conexão com o DB fechada")
	l.Info("serviço derrubado")
}

func getServerConfig(dbConfig *db.DbConfig, l *logger.Logger) *serverConfig {
	config := &serverConfig{
		logger:      l,
		dbConn:      db.InitDB(dbConfig, l),
		environment: env.Panic(env.Get("ENVIRONMENT_NAME").Required().String()),
	}

	portValue, err := env.Get("HTTP_PORT").Required().Uint(10, 64)
	if err != nil {
		l.Fatal("fatal! porta inválida especificada em HTTP_PORT", l.Prop("port", portValue))
	}
	config.portHTTP = uint16(portValue)

	return config
}

func getDBConfig() *db.DbConfig {
	var (
		user    = env.Panic(env.Get("DATABASE_USER").Required().String())
		passwd  = env.Panic(env.Get("DATABASE_PASSWD").Required().String())
		addr    = env.Panic(env.Get("DATABASE_ADDR").Required().String())
		port    = uint16(env.Panic(env.Get("DATABASE_PORT").Uint(10, 64)))
		dbName  = env.Panic(env.Get("DATABASE_NAME").Required().String())
		sslMode = env.Panic(env.Get("DATABASE_SSL_MODE").String())
	)

	return db.New(user, passwd, addr, dbName, sslMode, port)
}
