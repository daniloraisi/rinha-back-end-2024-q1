package rest

import (
	"net/http"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/daniloraisi/rinha-back-end/pkg/client"
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

func Transacionar(clientService client.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}
}
