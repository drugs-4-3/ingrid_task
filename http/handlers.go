package http

import (
	"log"
	"net/http"
)

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	if _, err := w.Write([]byte("not implemented")); err != nil {
		log.Println("error writing response")
	}
}
