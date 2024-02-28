package db

import (
	"context"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/jmoiron/sqlx"
)

type ClientService struct {
	Db       *sqlx.DB
	DbConfig *DbConfig
	Logger   *logger.Logger
}

func (c *ClientService) GetExtrato(ctx context.Context, id string) (interface{}, error) {
	c.Logger.Fatal("fatal! não implementado")
	return nil, nil
}

func (c *ClientService) Transacionar(ctx context.Context, id string, tx *sqlx.Tx) (interface{}, error) {
	c.Logger.Fatal("fatal! não implementado")
	return nil, nil
}
