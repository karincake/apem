package handlerhelper

import (
	"net/http"

	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"

	lg "github.com/karincake/apem/lang"
)

// Process error required for string
func requiredString(w http.ResponseWriter, fieldName, input string) bool {
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.XErrors{
			fieldName: te.XError{
				Code:    "val-required",
				Message: lg.I.Msg("val-required"),
			},
		}}, nil)
		return false
	}
	return true
}
