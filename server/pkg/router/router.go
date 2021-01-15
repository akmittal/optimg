package router

import (
	"github.com/go-chi/chi"
)

type Router struct {
	*chi.Mux
}

func Get() (*Router, error) {
	router, err := get()
	if err != nil {
		return nil, err
	}

	return &Router{
		router,
	}, nil
}

func get() (*chi.Mux, error) {
	router := chi.NewRouter()
	return router, nil
}
