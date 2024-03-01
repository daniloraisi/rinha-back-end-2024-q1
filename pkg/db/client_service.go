package db

import (
	"context"
	"database/sql"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	customErrors "github.com/daniloraisi/rinha-back-end/pkg/errors"
)

type ClientService struct {
	Db       *sql.DB
	DbConfig *DbConfig
	Logger   *logger.Logger
}

func (c *ClientService) GetExtrato(ctx context.Context, id string) (interface{}, error) {
	extrato := new(Extrato)
	switch err := c.Db.QueryRow("SELECT * FROM extrato($1)", id).Scan(&extrato); err {
	case sql.ErrNoRows:
		return nil, &customErrors.NotFound{Resource: "Extrato"}
	case nil:
		return &extrato, nil
	default:
		return nil, nil
	}
}

func (c *ClientService) Transacionar(ctx context.Context, id string, tx *sql.Tx) (interface{}, error) {
	c.Logger.Fatal("fatal! não implementado")
	return nil, nil
}
