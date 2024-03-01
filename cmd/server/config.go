package main

import (
	"database/sql"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
)

type serverConfig struct {
	portHTTP    uint16
	dbConn      *sql.DB
	logger      *logger.Logger
	environment string
}

func (config serverConfig) GetPortHTTP() uint16 {
	return config.portHTTP
}

func (config serverConfig) GetDbConn() *sql.DB {
	return config.dbConn
}

func (config serverConfig) GetLogger() *logger.Logger {
	return config.logger
}

func (config serverConfig) GetEnvironment() string {
	return config.environment
}
