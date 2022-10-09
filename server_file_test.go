package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writter http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writter, request, "./resources/ok.html")
	} else {
		http.ServeFile(writter, request, "./resources/notFound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// go:embed resources/ok.html
var resourceOk string

// go:embed resources/notFound.html
var resourceNotFound string

func ServeFileEmbed(writter http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writter, resourceOk)
	} else {
		fmt.Fprint(writter, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
