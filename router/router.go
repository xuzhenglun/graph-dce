package router

import (
	"net/http"
)

type Router struct {
	Warp      []func(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	Mux       map[string]http.HandlerFunc
	SubRouter []*Router
}

func NewRouter() *Router {
	return &Router{
		Mux: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) NewSubRouter() *Router {
	nr := NewRouter()
	r.SubRouter = append(r.SubRouter, nr)
	return nr
}

func (r *Router) Use(path string, handler http.HandlerFunc) {
	r.Mux[path] = handler
}

func (r *Router) UseFunc(warper func(func(http.ResponseWriter, *http.Request)) http.HandlerFunc) {
	r.Warp = append(r.Warp, warper)
}

func (r *Router) Register() {
	for path, handler := range r.Mux {
		for _, w := range r.Warp {
			handler = w(handler)
		}
		http.HandleFunc(path, handler)
	}

	if r.SubRouter != nil {
		for _, v := range r.SubRouter {
			v.Register()
		}
	}
}
