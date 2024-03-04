package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	customErrors "github.com/daniloraisi/rinha-back-end/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type ClientService struct {
	Db       *sqlx.DB
	DbConfig *DbConfig
	Logger   *logger.Logger
}

func (c *ClientService) GetExtrato(ctx context.Context, id string) (interface{}, error) {
	extrato := Extrato{}
	switch err := c.Db.Get(&extrato.Saldo, "SELECT s.valor, NOW()::TIMESTAMPTZ AS data_extrato, c.limite FROM saldos s JOIN clientes c USING(id) WHERE s.id_cliente = $1", id); err {
	case sql.ErrNoRows:
		return nil, &customErrors.NotFound{Resource: "Extrato"}
	case nil:
		err = c.Db.Select(&extrato.Transacoes, "SELECT valor, tipo, descricao, data_transacao FROM transacoes WHERE id_cliente = $1 ORDER BY id DESC LIMIT 10", id)
		if err != nil {
			return nil, &customErrors.InternalServerError{Err: err}
		}

		return &extrato, nil
	default:
		return nil, &customErrors.InternalServerError{Err: err}
	}
}

func (c *ClientService) Transacionar(ctx context.Context, id string, transacao TransacaoInput) (interface{}, error) {
	var tipo string
	switch strings.ToLower(transacao.Tipo) {
	case "d":
		tipo = "debito"
	case "c":
		tipo = "credito"
	default:
		return nil, &customErrors.TipoTransacaoInvalido{Reason: "Tipo inválido"}
	}

	novoSaldo := TransacaoOutput{}
	switch err := c.Db.Get(&novoSaldo, fmt.Sprintf("SELECT * FROM %s($1, $2, $3)", tipo), id, transacao.Valor, transacao.Descricao); err {
	case nil:
		if novoSaldo.ComErro {
			return nil, &customErrors.SaldoInsuficiente{Reason: novoSaldo.Mensagem}
		}

		return novoSaldo, nil
	default:
		return nil, &customErrors.InternalServerError{Err: err}
	}
}
