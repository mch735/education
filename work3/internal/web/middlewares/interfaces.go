package middlewares

import "net/http"

type Wrapper interface {
	HandlerFunc(h http.Handler) http.Handler
}
