package models

import (
	"net/http"
)

// ResponseContextWriter replaces http.ResponseWriter, to store response context information from service packages
type ResponseContextWriter struct {
	http.ResponseWriter
	Status int
	Length int
	Body   string

	// error info
	IsError      bool
	ErrorMessage string
	ClientError  string
}

func (w *ResponseContextWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *ResponseContextWriter) Write(b []byte) (int, error) {
	w.Body = string(b)
	n, err := w.ResponseWriter.Write(b)
	w.Length += n

	return n, err
}

func (w *ResponseContextWriter) WriteErrorJSON(status int, message string, b []byte) (int, error) {
	w.IsError = true
	w.ErrorMessage = message
	w.ClientError = string(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return w.Write(b)
}

func (w *ResponseContextWriter) WriteResponseJSON(status int, body []byte) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return w.Write(body)
}
