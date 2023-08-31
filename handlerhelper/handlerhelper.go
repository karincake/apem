package handlerhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	lz "github.com/karincake/apem/loggerzap"
	sv "github.com/karincake/serabi"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"

	lg "github.com/karincake/apem/lang"
)

// Writes json output thorugh http.ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) {
	js, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("{ \"message\": \"error converting data or result to json\"}"))
		w.WriteHeader(500)
		if lz.I != nil {
			lz.I.Error("error converting data or result to json")
		}
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

// Writes error, should be called by data submission
func WriteError(w http.ResponseWriter, err te.XErrors) {
	// check if the error is a non-field error
	if len(err) == 1 {
		for idx, httpCode := range ErrorCodes {
			if err.KeyExists(idx) {
				WriteJSON(w, httpCode, err[idx], nil)
				return
			}
		}
	}

	WriteJSON(w, http.StatusUnprocessableEntity, td.II{
		"meta":   td.IS{"count": strconv.Itoa(len(err))},
		"errors": err,
	}, nil)
}

// Respond at the service level that related to returning data
// Note that it should be called for things related to data processing due to
// it's only having 2 error conditions: data not found or unprocessable entity
func DataResponse(w http.ResponseWriter, meta, data, ref, err any) {
	if data == nil && err == nil {
		WriteJSON(w, http.StatusNotFound, te.XError{
			Code:    "data-notFound",
			Message: lg.I.Msg("data-notFound"),
		}, nil)
	} else if err != nil {
		if msgAsString, ok := err.(string); ok {
			WriteJSON(w, http.StatusUnprocessableEntity, td.IS{"Message": msgAsString}, nil)
		} else {
			if msgAsMap, ok := err.(te.XErrors); ok {
				WriteError(w, msgAsMap)
			} else if msgAsMap, ok := err.(te.XError); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, msgAsMap, nil)
			} else if msgAsMap, ok := err.(map[string]any); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{
					"meta":   td.IS{"count": strconv.Itoa(len(msgAsMap))},
					"errors": msgAsMap}, nil)
			} else if msgAsError, ok := err.(error); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, te.XError{
					Code:    "unknown",
					Message: msgAsError.Error(),
				}, nil)
			} else {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{"errors": err}, nil)
			}
		}
	} else {
		if message, ok := data.(string); ok {
			WriteJSON(w, http.StatusOK, td.IS{"message": message}, nil)
		} else {
			WriteJSON(w, http.StatusOK, td.Data{Meta: meta, Data: data}, nil)
		}
	}
}

// Validates a string assuming the field is required.
func ValidateString(w http.ResponseWriter, fieldName, input string) string {
	if !requiredString(w, fieldName, input) {
		return ""
	}
	return input
}

// Validates an int value from string, assuming the field is required.
func ValidateInt(w http.ResponseWriter, fieldName, input string) int {
	// val := chi.URLParam(r, input)
	if !requiredString(w, fieldName, input) {
		return 0
	}
	output, err := strconv.Atoi(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.XErrors{
			fieldName: te.XError{
				Code:    "val-int",
				Message: lg.I.Msg("val-int"),
			},
		}}, nil)
		return 0
	}
	return output
}

// Validates a UUID from string assuming the field is required.
func ValidateIdUuid(w http.ResponseWriter, fieldName, input string) uuid.UUID {
	if !requiredString(w, fieldName, input) {
		return uuid.Nil
	}
	output, err := uuid.Parse(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.XErrors{
			fieldName: te.XError{
				Code:    "val-uuid",
				Message: lg.I.Msg("val-uuid"),
			},
		}}, nil)
		return uuid.Nil
	}
	return output
}

// Validates struct
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		WriteError(w, err.(te.XErrors))
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		WriteError(w, err.(te.XErrors))
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		WriteError(w, err.(te.XErrors))
		return false
	}

	return true
}

// by form-data version of ValidateStruct, to cover form-data, return bool true on success
func ValidateStructByFD(w http.ResponseWriter, r *http.Request, data any) bool {
	err := sv.ValidateFormData(&data, r)
	if err != nil {
		WriteError(w, err.(te.XErrors))
		return false
	}

	teErr := sv.Validate(data)
	if teErr != nil {
		WriteError(w, teErr.(te.XErrors))
		return false
	}

	return true
}
