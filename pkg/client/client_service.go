package client

import (
	"context"

	"github.com/daniloraisi/rinha-back-end/pkg/db"
)

type Client interface {
	GetExtrato(ctx context.Context, id string) (interface{}, error)
	Transacionar(ctx context.Context, id string, transacao db.TransacaoInput) (interface{}, error)
}
