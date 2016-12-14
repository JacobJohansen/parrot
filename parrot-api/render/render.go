package render

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var (
	jsonContentType        = "application/json"
	jsonContentTypeCharset = "application/json; charset=utf-8"
)

type apiResponseBody struct {
	responseMeta `json:"meta,omitempty"`
	Payload      interface{} `json:"payload,omitempty"`
}

type responseMeta struct {
	Status int   `json:"status,omitempty"`
	Error  error `json:"error,omitempty"`
}

func Error(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(status)

	body := apiResponseBody{
		responseMeta: responseMeta{
			Status: status,
			Error:  err},
		Payload: nil}

	handleBody(w, body)
}

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(status)

	body := apiResponseBody{
		responseMeta: responseMeta{
			Status: status},
		Payload: payload}

	handleBody(w, body)
}

func JSONWithHeaders(w http.ResponseWriter, status int, headers map[string]string, payload interface{}) {
	h := w.Header()
	h.Set("Content-Type", jsonContentType)
	for k, v := range headers {
		h.Set(k, v)
	}
	w.WriteHeader(status)

	body := apiResponseBody{
		responseMeta: responseMeta{
			Status: status},
		Payload: payload}

	handleBody(w, body)
}

func handleBody(w http.ResponseWriter, body apiResponseBody) {
	encoded, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		logrus.Error(err)
	}

	w.Write(encoded)
}