package handlerhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	lz "github.com/karincake/apem/loggerzap"
	sv "github.com/karincake/serabi"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"

	lg "github.com/karincake/apem/lang"
)

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

// write error response if validation fails, return boool true if success
func ValidateAutoInc(w http.ResponseWriter, r *http.Request, input string) int {
	output, err := strconv.Atoi(chi.URLParam(r, input))
	if err != nil || output < 1 {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(input, te.NewError("val-integerPositive", lg.I.Msg("val-integerPositive")))}, nil)
		return 0
	}
	return output
}

func ValidateString(w http.ResponseWriter, r *http.Request, input string) string {
	output := chi.URLParam(r, input)
	if output == "" {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(input, te.NewError("val-required", lg.I.Msg("val-required")))}, nil)
		return ""
	}
	return output
}

// write error if
func ValidateIdUuid(w http.ResponseWriter, input string) uuid.UUID {
	output, err := uuid.Parse(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, td.II{"errors": te.NewErrorsPick(input, te.NewError("val-validUuid", lg.I.Msg("val-validUuid")))}, nil)
		return uuid.Nil
	}
	return output
}

// write error response if validation fails, return boool true on success
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		writeError(w, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		writeError(w, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		writeError(w, err)
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

// internal
func writeError(w http.ResponseWriter, err te.Errors) {
	if err.KeyExists("struct") {
		WriteJSON(w, http.StatusBadRequest, err.GetOne("struct"), nil)
	} else if err.KeyExists("resource-notFound") {
		WriteJSON(w, http.StatusNotFound, err.GetOne("resource-notFound"), nil)
	} else {
		WriteJSON(w, http.StatusUnprocessableEntity, td.II{
			"meta":   td.IS{"count": strconv.Itoa(err.Count())},
			"errors": err,
		}, nil)
	}
}
