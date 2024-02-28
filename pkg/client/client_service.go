package client

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	GetExtrato(ctx context.Context, id string) (interface{}, error)
	Transacionar(ctx context.Context, id string, tx *sqlx.Tx) (interface{}, error)
}
