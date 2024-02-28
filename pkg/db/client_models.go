package db

import "time"

type Saldo struct {
	Total       uint64    `json:"saldo"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      uint64    `json:"limite"`
}

type Transacao struct {
	Valor       uint64    `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type Extrato struct {
	Saldo      Saldo       `json:"saldo"`
	Transacoes []Transacao `json:"ultimas_transacoes"`
}
