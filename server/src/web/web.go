package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Start() error {
	r := chi.NewRouter()
	r.Post("/optimize", Optimize)
	r.Get("/gallery", Gallery)
	r.Get("/img", ImageHandler)
	r.Handle("/images/src/*", http.StripPrefix("/images/src/", http.FileServer(http.Dir("/Users/amittal/projects/images"))))
	r.Handle("/images/target/*", http.StripPrefix("/images/target/", http.FileServer(http.Dir("/Users/amittal/images"))))
	imageRouter := chi.NewRouter()
	imageRouter.HandleFunc("/*", ImageServer)
	go http.ListenAndServe(":4000", imageRouter)

	return http.ListenAndServe(":3000", r)
}
