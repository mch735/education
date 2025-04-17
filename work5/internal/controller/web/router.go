package web

import (
	"net/http"

	"github.com/mch735/education/work5/internal/usecase"
)

func NewRouter(uc usecase.User) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/users", Index(uc))
	mux.HandleFunc("GET /v1/users/{id}", Show(uc))
	mux.HandleFunc("POST /v1/users", Create(uc))
	mux.HandleFunc("PUT /v1/users/{id}", Update(uc))
	mux.HandleFunc("DELETE /v1/users/{id}", Delete(uc))

	return mux
}
