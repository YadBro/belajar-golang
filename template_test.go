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

func SimpleHTML(writter http.ResponseWriter, request *http.Request) {
	templateText := `<html><head><body>{{.}}</body></head></html>`

	// Cara Manual
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writter, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTMLTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLTemplateFile(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML Template File")
}
func TestSimpleHTMLTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLTemplateFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML Template File")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello HTML Template File")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TEMPLATE DATA MAP
func TemplateDataMap(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello World")
	t.ExecuteTemplate(writter, "name.gohtml", map[string]interface{}{
		"Title":    "Template Data Map",
		"FullName": "Yadi Apriyadi",
		"Address": map[string]string{
			"Street": "Jalan map",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title    string
	FullName string
	Address
}

// TEMPLATE DATA Struct
func TemplateDataStruct(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "simple.gohtml", "Hello World")
	t.ExecuteTemplate(writter, "name.gohtml", Page{
		Title:    "Template Data Struct",
		FullName: "Yadi Apriyadi",
		Address: Address{
			Street: "Jalan struct",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TEMPLATE ACTION
func TemplateActionIf(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	// t.ExecuteTemplate(writter, "if.gohtml", map[string]string{
	// 	"Title": "Template IF",  // FullName Not Found!
	// })
	// t.ExecuteTemplate(writter, "if.gohtml", map[string]string{
	// 	"Title":    "Template IF",
	// 	"FullName": "Yadi Apriyadi",
	// })
	t.ExecuteTemplate(writter, "if.gohtml", map[string]string{
		"Title":    "Template IF",
		"FullName": "udin",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

func TemplateActionRange(writter http.ResponseWriter, request *http.Request) {
	addFuncs := template.FuncMap{"add": add}
	t := template.Must(template.New("range.gohtml").Funcs(addFuncs).ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writter, "range.gohtml", map[string]interface{}{
		"Title":   "Template RANGE",
		"Hobbies": []string{"Soccer", "Cooking", "FootBall", "BasketBall", "Game", "Read", "Code"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

func add(x, y int) int {
	return x + y
}

func TestServerTemplateAction(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/hobbies", TemplateActionRange)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func TemplateActionWith(writter http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/address.gohtml"))
	// t.ExecuteTemplate(writter, "address.gohtml", map[string]interface{}{
	// 	"Title": "Template WITH",
	// })

	t.ExecuteTemplate(writter, "address.gohtml", map[string]interface{}{
		"Title": "Template WITH",
		"Address": map[string]string{
			"Street": "Jalan Kemayoran Lama",
			"City":   "Cimahi",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
