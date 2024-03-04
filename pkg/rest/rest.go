package rest

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/daniloraisi/rinha-back-end/internal/logger"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func ResetDB(db *sqlx.DB, l *logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		_, err := db.Exec("SELECT reset_db()")
		if err != nil {
			ErrorJSON(w, err, l)
			return
		}

		l.Info("reset-db! Executado com sucesso.")
		ResponseJSON(w, "ok", l)
	}
}

func ResponseJSON(w http.ResponseWriter, response interface{}, l *logger.Logger) {
	ResponseJSONStatus(w, http.StatusOK, &response, l)
}

func ResponseJSONStatus(w http.ResponseWriter, statusCode int, response interface{}, l *logger.Logger) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if response == nil || (reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil()) {
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		l.Error("error!", l.ErrProp(err))
	}
}

func ErrorJSON(w http.ResponseWriter, err error, l *logger.Logger) {
	ErrorJSONStatus(w, http.StatusInternalServerError, err, l)
}

func ErrorJSONStatus(w http.ResponseWriter, statusCode int, err error, l *logger.Logger) {
	response := map[string]map[string]string{"error": {"message": err.Error()}}

	ResponseJSONStatus(w, statusCode, response, l)
}

func ParseJSON[T interface{}](w http.ResponseWriter, r *http.Request, model T) error {
	return json.NewDecoder(r.Body).Decode(&model)
}
