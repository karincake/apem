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

	fh "github.com/karincake/apem/formdatahelper"
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
func WriteError(w http.ResponseWriter, err te.Errors) {
	// check if the error is a non-field error
	if err.Count() == 1 {
		for idx, httpCode := range ErrorCodes {
			if err.KeyExists(idx) {
				WriteJSON(w, httpCode, err.GetOne(idx), nil)
				return
			}
		}
	}

	WriteJSON(w, http.StatusUnprocessableEntity, td.II{
		"meta":   td.IS{"count": strconv.Itoa(err.Count())},
		"errors": err,
	}, nil)
}

// respond at the service level that related to returning data
func DataResponse(w http.ResponseWriter, meta, data, ref, err any) {
	if data == nil && err == nil {
		WriteJSON(w, http.StatusNotFound, te.NewError("data-notFound", lg.I.Msg("data-notFound")), nil)
	} else if err != nil {
		if msgAsString, ok := err.(string); ok {
			WriteJSON(w, http.StatusUnprocessableEntity, td.IS{"Message": msgAsString}, nil)
		} else {
			if msgAsMap, ok := err.(te.Errors); ok {
				WriteError(w, msgAsMap)
			} else if msgAsMap, ok := err.(map[string]any); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{
					"meta":   td.IS{"count": strconv.Itoa(len(msgAsMap))},
					"errors": err}, nil)
			} else if msgAsError, ok := err.(error); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, te.NewError("unknown", msgAsError.Error()), nil)
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
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(fieldName, te.NewError("val-required", lg.I.Msg("val-required")))}, nil)
		return ""
	}
	return input
}

// Validates an int value from string, assuming the field is required.
func ValidateInt(w http.ResponseWriter, fieldName, input string) int {
	// val := chi.URLParam(r, input)
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(fieldName, te.NewError("val-required", lg.I.Msg("val-required")))}, nil)
		return 0
	}
	output, err := strconv.Atoi(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(fieldName, te.NewError("val-int", lg.I.Msg("val-int")))}, nil)
		return 0
	}
	return output
}

// Validates a UUID from string assuming the field is required.
func ValidateIdUuid(w http.ResponseWriter, fieldName, input string) uuid.UUID {
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(fieldName, te.NewError("val-required", lg.I.Msg("val-required")))}, nil)
		return uuid.Nil
	}
	output, err := uuid.Parse(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(fieldName, te.NewError("val-uuid", lg.I.Msg("val-uuid")))}, nil)
		return uuid.Nil
	}
	return output
}

// Validates struct
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		WriteError(w, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		WriteError(w, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		WriteError(w, err)
		return false
	}

	return true
}

// by form-data version of ValidateStruct, to cover form-data, return bool true on success
func ValidateStructByFD(w http.ResponseWriter, r *http.Request, data any) bool {
	err := fh.CopyToStruct(&data, r)
	if err != nil {
		errors := te.NewErrorsPick("payload-bad", te.NewError("parse-fail", err.Error()))
		WriteError(w, errors)
		return false
	}

	return true
}
