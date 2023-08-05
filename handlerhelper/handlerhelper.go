package handlerhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	sv "github.com/karincake/serabi"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"

	lg "github.com/karincake/apem/lang"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ProcessService(w http.ResponseWriter, r *http.Request, result, err any) {
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, td.II{"errors": err}, nil)
	} else {
		DataResponse(w, nil, result, nil, nil)
	}
}

// write error response if validation fails, return boool true if success
func ValidateAutoInc(w http.ResponseWriter, r *http.Request, input string) int {
	id, err := strconv.Atoi(chi.URLParam(r, input))
	if err != nil || id < 1 {
		WriteJSON(w, http.StatusBadRequest, te.NewError("fieldVal-integerPositive", lg.I.Msg("fieldVal-integerPositive")), nil)
		return 0
	}
	return id
}

func ValidateString(w http.ResponseWriter, r *http.Request, input string) string {
	result := chi.URLParam(r, input)
	if result == "" {
		WriteJSON(w, http.StatusBadRequest, te.NewError("field-required", lg.I.Msg("field-required")), nil)
		return ""
	}
	return result
}

// write error if
func ValidateIdUuid(w http.ResponseWriter, id string) (uid uuid.UUID, pass bool) {
	uid, err := uuid.Parse(id)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, te.NewError("field-validUuid", lg.I.Msg("field-validUuid")+", "+err.Error()), nil)
		pass = false
		return
	}

	pass = true
	return
}

// write error response if validation fails, return boool true on success
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if err.KeyExists("struct") {
			httpStatus = http.StatusBadRequest
		}
		WriteJSON(w, httpStatus, td.II{
			"meta":   td.IS{"count": strconv.Itoa(err.Count())},
			"errors": err,
		}, nil)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if err.KeyExists("struct") {
			httpStatus = http.StatusBadRequest
		}
		WriteJSON(w, httpStatus, td.II{
			"meta":   td.IS{"count": strconv.Itoa(err.Count())},
			"errors": err,
		}, nil)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if err.KeyExists("struct") {
			httpStatus = http.StatusBadRequest
		}
		WriteJSON(w, httpStatus, td.II{
			"meta":   td.IS{"count": strconv.Itoa(err.Count())},
			"errors": err,
		}, nil)
		return false
	}

	return true
}

// respond at the service level that related to returning data
func DataResponse(w http.ResponseWriter, meta, data, ref, err any) {
	if data == nil && err == nil {
		WriteJSON(w, http.StatusNotFound, te.NewError("resource-notFound", lg.I.Msg("resource-notFound")), nil)
	} else if err != nil {
		if msgAsString, ok := err.(string); ok {
			WriteJSON(w, http.StatusUnprocessableEntity, td.IS{"Message": msgAsString}, nil)
		} else {
			if msgAsMap, ok := err.(te.Errors); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, td.II{
					"meta":   td.IS{"count": strconv.Itoa(msgAsMap.Count())},
					"errors": msgAsMap}, nil)
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
		if message, ok := data.(string); !ok {
			WriteJSON(w, http.StatusOK, td.Data{Meta: meta, Data: data}, nil)
		} else {
			WriteJSON(w, http.StatusOK, td.IS{"message": message}, nil)
		}
	}
}
