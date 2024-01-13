package handlerhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/google/uuid"
	lz "github.com/karincake/apem/loggerzap"
	sv "github.com/karincake/serabi"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"

	lg "github.com/karincake/apem/lang"
)

// The primary function that writes json output through http.ResponseWriter
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

// Writes an error response, will try to match list of recognized error to decide
// the http status code
func WriteError(w http.ResponseWriter, err te.XError) {
	httpCode := http.StatusUnprocessableEntity
	for idx, code := range ErrorCodes {
		if err.Code == idx {
			httpCode = code
			break
		}
	}

	WriteJSON(w, httpCode, err, nil)
}

// Generate response based on the condition of data and error
// Note that it should be called for things related to data processing and
// with the possible error only for unprocessable entity that related to
// data field which can have more than one error. Any non-field error, which
// normally a single error will be disposed to WriteError funcion
func DataResponse(w http.ResponseWriter, data, err any) {
	if data != nil && err == nil {
		if dataVal, ok := data.(td.Data); ok {
			WriteJSON(w, http.StatusOK, dataVal, nil)
		} else if message, ok := data.(string); ok {
			WriteJSON(w, http.StatusOK, td.IS{"message": message}, nil)
		} else {
			v := reflect.ValueOf(data)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			vKind := v.Kind()
			if vKind != reflect.Struct && vKind != reflect.Map {
				WriteJSON(w, http.StatusOK, td.II{"value": data}, nil)
			} else {
				WriteJSON(w, http.StatusOK, data, nil)
			}
		}
	} else if err != nil {
		if stringErr, ok := err.(string); ok {
			WriteJSON(w, http.StatusUnprocessableEntity, td.IS{"Message": stringErr}, nil)
		} else {
			if castedErr, ok := err.(te.XError); ok {
				// this is the only error that can have non unprocessable entity error
				WriteError(w, castedErr)
			} else if castedErr, ok := err.(te.XErrors); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{
					"meta":   td.IS{"count": strconv.Itoa(len(castedErr))},
					"errors": err,
				}, nil)
			} else if castedErr, ok := err.(map[string]any); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{
					"meta":   td.IS{"count": strconv.Itoa(len(castedErr))},
					"errors": castedErr}, nil)
			} else if castedErr, ok := err.(error); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, te.XError{
					Code:    "unknown",
					Message: castedErr.Error(),
				}, nil)
			} else {
				// worst case unknown error
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{"errors": err}, nil)
			}
		}
	} else {
		WriteJSON(w, http.StatusNotFound, te.XError{
			Code:    "data-notFound",
			Message: lg.I.Msg("data-notFound"),
		}, nil)
	}
}

// Validates a string assuming the field is required.
// Error occured will be categorized as data-field
func ValidateString(w http.ResponseWriter, fieldName, input string) string {
	if !requiredString(w, fieldName, input) {
		return ""
	}
	return input
}

// Validates an int value from string, assuming the field is required.
// Error occured will be categorized as data-field
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

// Validates a UUID from string assuming the field is required
// Error occured will be categorized as data-field
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
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by form-data version of ValidateStruct, to cover form-data, return bool true on success
func ValidateStructByFD(w http.ResponseWriter, r *http.Request, data any) bool {
	err := sv.ValidateFormData(&data, r)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}
