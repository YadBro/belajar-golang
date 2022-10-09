package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatesCaching embed.FS

var myTemplate = template.Must(template.ParseFS(templatesCaching, "templates/simple.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML TemplateCaching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
