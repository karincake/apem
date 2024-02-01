package apem

import (
	"encoding/json"
	"net/http"

	lz "github.com/karincake/apem/loggerzap"
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
