package render

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func ErrorJSON(w http.ResponseWriter, err error, status int) {
	jSON(w, response{
		Success: false,
		Status: status,
		Error:  err.Error(),
		Data:   []interface{}{},
	})
}

func SingleSuccessJSON(w http.ResponseWriter, v interface{}) {
	jSON(w, response{
		Success: true,
		Status: http.StatusOK,
		Data:   []interface{}{v},
	})
}

func MultipleSuccessJSON(w http.ResponseWriter, v interface{}) {
	jSON(w, response{
		Success: true,
		Status: http.StatusOK,
		Data:   v,
	})
}

func jSON(w http.ResponseWriter, v interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buf.Bytes())

}
