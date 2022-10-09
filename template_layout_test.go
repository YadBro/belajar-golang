package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/header.gohtml", "templates/layout.gohtml", "templates/footer.gohtml"))

	t.ExecuteTemplate(writter, "layout", map[string]interface{}{
		"Title": "Template LAYOUT",
		"Name":  "Yadi",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
