package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my Name is " + myPage.Name
}

func TemplateFunction(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}} {{len .Name}}`))

	t.ExecuteTemplate(writter, "FUNCTION", MyPage{
		Name: "Yadi",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

func TemplateFunctionGlobal(writter http.ResponseWriter, request *http.Request) {
	t := template.New("GLOBAL FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	t = template.Must(t.Parse("{{ToUpper .Name}}"))

	t.ExecuteTemplate(writter, "GLOBAL FUNCTION", MyPage{
		Name: "Yadi",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

func TemplateFunctionGlobalPipelines(writter http.ResponseWriter, request *http.Request) {
	t := template.New("GLOBAL FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"SayHello": func(name string) string {
			return "Hello " + name
		},
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	t = template.Must(t.Parse("{{SayHello .Name | ToUpper}}"))

	t.ExecuteTemplate(writter, "GLOBAL FUNCTION", MyPage{
		Name: "Yadi",
	})
}

func TestTemplateFunctionGlobalPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalPipelines(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
