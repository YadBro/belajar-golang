package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

type ErrorMiddleware struct {
	http.Handler
}

func (middleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()

		fmt.Println("RECOVER:", err)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Ada error yaitu: %s", err)
		}
	}()
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprintf(writer, "Hello Middleware")
	})

	// Error Handler
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		panic("Ups")
	})

	// Log Middleware (Handler Pertama)
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// Error LogMiddleware (Handler Kedua)
	errorMiddleWare := &ErrorMiddleware{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorMiddleWare,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
