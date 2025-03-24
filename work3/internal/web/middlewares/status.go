package middlewares

import (
	"net/http"
)

type StatusHandler struct {
	ProcessingFunc func(method string, path string, statusCode int, statusText string)
}

func (f StatusHandler) HandlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		sw := StatusWriter{w, http.StatusOK}

		h.ServeHTTP(&sw, req)

		f.ProcessingFunc(req.Method, req.URL.Path, sw.StatusCode, http.StatusText(sw.StatusCode))
	})
}

type StatusWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *StatusWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
