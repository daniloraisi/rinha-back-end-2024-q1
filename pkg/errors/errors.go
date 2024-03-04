package errors

import (
	"fmt"
	"os"
)

type BadRequest struct {
	Reason string
}

func (n *BadRequest) Error() string {
	if n.Reason != "" {
		return "Bad Request: " + n.Reason
	}

	return "Bad Request"
}

type NotFound struct {
	Resource string
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("%s not found", n.Resource)
}

type InvalidBody struct {
	Reason string
}

func (n *InvalidBody) Error() string {
	if n.Reason != "" {
		return "Invalid body: " + n.Reason
	}

	return "Invalid body"
}

type InternalServerError struct {
	Err error
}

func (i *InternalServerError) Error() string {
	fmt.Fprintln(os.Stderr, i.Err)
	return "Internal server error"
}

type SaldoInsuficiente struct {
	Reason string
}

func (si *SaldoInsuficiente) Error() string {
	return si.Reason
}

type TipoTransacaoInvalido struct {
	Reason string
}

func (tti *TipoTransacaoInvalido) Error() string {
	return tti.Reason
}
