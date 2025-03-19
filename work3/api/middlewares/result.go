package middlewares

import (
	"net/http"
	"strconv"
)

type ResultHandler struct{}

func (r ResultHandler) HandlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		code := req.URL.Query().Get("code")

		statusCode, err := strconv.Atoi(code)
		if err == nil && http.StatusText(statusCode) != "" {
			w.WriteHeader(statusCode)

			return
		}

		h.ServeHTTP(w, req)
	})
}
