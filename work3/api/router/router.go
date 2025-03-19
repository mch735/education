package router

import (
	"net/http"

	"github.com/mch735/education/work3/api/middlewares"
)

type Router struct {
	middlewares []middlewares.Wrapper
	http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		make([]middlewares.Wrapper, 0),
		http.ServeMux{},
	}
}

func (r *Router) Middleware(wrapper middlewares.Wrapper) {
	r.middlewares = append(r.middlewares, wrapper)
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h := http.Handler(http.HandlerFunc(handler))

	for _, wrapper := range r.middlewares {
		h = wrapper.HandlerFunc(h)
	}

	r.Handle(pattern, h)
}
