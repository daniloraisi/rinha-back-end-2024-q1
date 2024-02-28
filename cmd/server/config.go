package main

import (
	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/jmoiron/sqlx"
)

type serverConfig struct {
	portHTTP    uint16
	dbConn      *sqlx.DB
	logger      *logger.Logger
	environment string
}

func (config serverConfig) GetPortHTTP() uint16 {
	return config.portHTTP
}

func (config serverConfig) GetDbConn() *sqlx.DB {
	return config.dbConn
}

func (config serverConfig) GetLogger() *logger.Logger {
	return config.logger
}

func (config serverConfig) GetEnvironment() string {
	return config.environment
}
