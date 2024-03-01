package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Saldo struct {
	Total       uint64    `json:"total"`
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

func (e Extrato) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e *Extrato) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &e)
}

func (t Transacao) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Transacao) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &t)
}
