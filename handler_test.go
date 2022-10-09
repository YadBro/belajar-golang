package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		_, err := fmt.Fprint(w, "<h1 style='color: red; border: 1px solid black;'>Hello Yadi</h1>")
		if err != nil {
			panic(err)
		}

	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "<h1 style='color: red; border: 1px solid black;'>Hello World, My name is Yadi Apriyadi</h1>")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "<h1 style='color: red; border: 1px solid black;'>I'am a programmer</h1>")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/social-media", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "<h1 style='color: red; border: 1px solid black;'>Follow my <a href='https://www.linkedin.com/in/yadi-apriyadi/' target='_blank' rel='noreferrer'>LinkedIn</a></h1>")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Image")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Thumbnail")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		_, err := fmt.Fprint(w, "<h1 style='color: red; border: 1px solid black;'>Hello Yadi</h1>")
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		if err != nil {
			panic(err)
		}

	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
