package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Start() error {
	r := chi.NewRouter()
	r.Post("/api/optimize", Optimize)
	r.Get("/api/gallery", Gallery)
	r.Get("/api/img", ImageHandler)
	r.Handle("/api/images/src/*", http.StripPrefix("/api/images/src/", http.FileServer(http.Dir(sourcePATH))))
	r.Handle("/api/images/target/*", http.StripPrefix("/api/images/target/", http.FileServer(http.Dir(targetPath))))
	imageRouter := chi.NewRouter()
	imageRouter.HandleFunc("/api/*", ImageServer)
	go http.ListenAndServe(":8001", imageRouter)

	return http.ListenAndServe(":8000", r)
}
