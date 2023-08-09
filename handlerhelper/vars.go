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

var ErrorCodes = map[string]int{
	"payload-bad":    http.StatusBadRequest,
	"data-notFound":  http.StatusNotFound,
	"auth-forbidden": http.StatusForbidden,
}
