package render

import (
	"bytes"
	"encoding/json"
	"github.com/serhatmorkoc/go-realworld-example/handler/api/errors"
	"net/http"
)

func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &errors.Error{Message: err.Error()}, status)
}

func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusNotFound)
}

func Unauthorized(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusForbidden)
}

func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusBadRequest)
}

func JSON(w http.ResponseWriter, v interface{}, status int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err := w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
