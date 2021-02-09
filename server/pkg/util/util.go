package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Msg   string `json:"msg"`
	Field string `json:"field"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func HTTPError(rw http.ResponseWriter, data Error, code int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}
func HTTPJson(rw http.ResponseWriter, data interface{}) {

	json.NewEncoder(rw).Encode(data)
}
