package image

import (
	"net/http"
)

func HandleStaticSource() http.Handler {

	fs := http.FileServer(http.Dir("/Users/amittal/images/source"))
	return http.StripPrefix("/api/static/source/", fs)
}
func HandleStaticDest() http.Handler {

	fs := http.FileServer(http.Dir("/Users/amittal/images/dest"))
	return http.StripPrefix("/api/static/dest/", fs)
}
