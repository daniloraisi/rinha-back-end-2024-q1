package db

import (
	"time"
)

type TransacaoInput struct {
	Valor     uint64 `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransacaoOutput struct {
	NovoSaldo int64  `db:"novo_saldo" json:"saldo"`
	Limite    int64  `db:"limite" json:"limite"`
	ComErro   bool   `db:"com_erro" json:"-"`
	Mensagem  string `db:"mensagem" json:"-"`
}

type Saldo struct {
	Total       int64     `json:"total" db:"valor"`
	DataExtrato time.Time `json:"data_extrato" db:"data_extrato"`
	Limite      int64     `json:"limite" db:"limite"`
}

type Transacao struct {
	Valor       int64     `json:"valor" db:"valor"`
	Tipo        string    `json:"tipo" db:"tipo"`
	Descricao   string    `json:"descricao" db:"descricao"`
	RealizadaEm time.Time `json:"realizada_em" db:"data_transacao"`
}

type Extrato struct {
	Saldo      Saldo       `json:"saldo"`
	Transacoes []Transacao `json:"ultimas_transacoes,omitempty"`
}
