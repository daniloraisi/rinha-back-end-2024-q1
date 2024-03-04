package rest

import (
	"errors"
	"net/http"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/pkg/client"
	"github.com/daniloraisi/rinha-back-end/pkg/db"
	customErrors "github.com/daniloraisi/rinha-back-end/pkg/errors"
	"github.com/julienschmidt/httprouter"
)

func GetExtrato(clientService client.Client, l *logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		extrato, err := clientService.GetExtrato(r.Context(), p.ByName("id"))
		if err != nil {
			switch err.(type) {
			case *customErrors.BadRequest:
				ErrorJSONStatus(w, http.StatusBadRequest, err, l)
				return
			case *customErrors.NotFound:
				ErrorJSONStatus(w, http.StatusNotFound, err, l)
				return
			default:
				ErrorJSON(w, err, l)
				return
			}
		}

		ResponseJSONStatus(w, http.StatusOK, extrato, l)
	}
}

func Transacionar(clientService client.Client, l *logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var body db.TransacaoInput
		if err := ParseJSON(w, r, &body); err != nil {
			ErrorJSONStatus(w, http.StatusBadRequest, &customErrors.InvalidBody{}, l)
			return
		}

		if len(body.Descricao) == 0 || len(body.Descricao) > 10 {
			ErrorJSONStatus(w, http.StatusUnprocessableEntity, errors.New("Descricão inválida"), l)
			return
		}

		transacao, err := clientService.Transacionar(r.Context(), p.ByName("id"), body)
		if err != nil {
			switch err.(type) {
			case *customErrors.SaldoInsuficiente:
				ErrorJSONStatus(w, http.StatusUnprocessableEntity, err, l)
				return
			case *customErrors.TipoTransacaoInvalido:
				ErrorJSONStatus(w, http.StatusUnprocessableEntity, err, l)
				return
			default:
				ErrorJSON(w, err, l)
				return
			}
		}

		ResponseJSON(w, transacao, l)
	}
}
