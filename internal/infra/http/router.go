package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/foorester/cook/internal/sys"
)

type (
	Router interface {
		chi.Router
		Method(method, pattern string, handler http.Handler)
		Mount(pattern string, handler http.Handler)
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	SimpleRouter struct {
		sys.Worker
		chi.Router
	}
)

func NewRouter(name string, opts ...sys.Option) Router {
	return &SimpleRouter{
		Worker: sys.NewWorker(name, opts...),
		Router: chi.NewRouter(),
	}
}
