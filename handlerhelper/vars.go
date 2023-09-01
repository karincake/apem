package handlerhelper

import (
	"net/http"
)

// Note that all the codes here are for data processing result
// that is used by the handler helper. The codes level are in not
// in the data field level.

var SuccessCodes = map[string]int{
	"request-ok":    http.StatusOK,
	"data-created":  http.StatusCreated,
	"data-accepted": http.StatusAccepted,
}

// mostly for single error
var ErrorCodes = map[string]int{
	"auth-required":  http.StatusUnauthorized,
	"auth-forbidden": http.StatusForbidden,
	"data-notFound":  http.StatusNotFound,
	"server-error":   http.StatusInternalServerError,
	"payload-bad":    http.StatusBadRequest,
	"parse-fail":     http.StatusBadRequest,
	"convert-fail":   http.StatusBadRequest,
}
