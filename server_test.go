package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	var handler2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprint(w, "hello")
		} else {
			fmt.Fprintf(w, "Hello %s", name)
		}
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler2,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
