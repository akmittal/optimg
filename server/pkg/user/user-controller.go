package user

import (
	"fmt"
	"net/http"
)

func UserController() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "%s", "Hello world")
	}
}
