package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Start() error {
	r := chi.NewRouter()
	r.Post("/", Optimize)
	imageRouter := chi.NewRouter()
	imageRouter.HandleFunc("/*", ImageServer)
	go http.ListenAndServe(":4000", imageRouter)

	return http.ListenAndServe(":3000", r)
}
