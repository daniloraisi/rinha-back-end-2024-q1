package db

import (
	"time"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/internal/pgsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	User    string
	Passwd  string
	Addr    string
	Port    uint16
	DBName  string
	SSLMode string
}

func New(user, passwd, addr, dbName, sslMode string, port uint16) *DbConfig {
	return &DbConfig{
		User:    user,
		Passwd:  passwd,
		Addr:    addr,
		Port:    port,
		DBName:  dbName,
		SSLMode: sslMode,
	}
}

func InitDB(config *DbConfig, l *logger.Logger) *sqlx.DB {
	dbConfig := &pgsql.Config{
		User:    config.User,
		Passwd:  config.Passwd,
		Addr:    config.Addr,
		DBName:  config.DBName,
		SSLMode: config.SSLMode,
	}
	l.Debug(dbConfig.FormatDSN())
	db, err := sqlx.Open("postgres", dbConfig.FormatDSN())
	if err != nil {
		l.Fatal("fatal! não foi possível inicializar o DB", l.ErrProp(err))
	}

	err = db.Ping()
	if err != nil {
		l.Fatal("fatal! erro de conexão com o DB", l.ErrProp(err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
