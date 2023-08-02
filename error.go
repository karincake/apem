package apem

import (
	"net/http"

	td "github.com/karincake/tempe/data"
	"go.uber.org/zap"

	hh "github.com/karincake/apem/handlerhelper"
	lz "github.com/karincake/apem/loggerzap"
)

func (a *app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := td.II{"error": message}
	err := hh.WriteJSON(w, status, env, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(500)
	}
}

func (a *app) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	a.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *app) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *app) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	a.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (a *app) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	a.errorResponse(w, r, http.StatusConflict, message)
}

func (a *app) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	a.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (a *app) logError(r *http.Request, err error) {
	lz.I.Info(err.Error(), zap.String("request_method", r.Method), zap.String("request_url", r.URL.String()))
}
