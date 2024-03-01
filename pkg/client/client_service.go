package client

import (
	"context"
	"database/sql"
)

type Client interface {
	GetExtrato(ctx context.Context, id string) (interface{}, error)
	Transacionar(ctx context.Context, id string, tx *sql.Tx) (interface{}, error)
}
