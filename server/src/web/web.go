package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Start() error {
	r := chi.NewRouter()
	r.Post("/", Optimize)
	return http.ListenAndServe(":3000", r)
}
