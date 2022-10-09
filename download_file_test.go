package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	image := request.URL.Query().Get("image")

	if image == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request!")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+image+"\"")
	http.ServeFile(writer, request, "./resources/"+image)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
