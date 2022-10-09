package belajar_golang_web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/*.gohtml
var templatesCaching3 embed.FS

var myTemplate3 = template.Must(template.ParseFS(templatesCaching3, "templates/*.gohtml"))

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplate3.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// Di balik FormFile dia ngelakuin ini dulu, default upload file di golang 32mb kita bisa ganti dengan ParseMultipartForm
	// request.ParseMultipartForm(100 << 20) // 100mb

	file, fileHeader, err := request.FormFile("image") // input name attribute filenya
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	myTemplate3.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	serverUpload := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := serverUpload.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/maskot.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	// Setup
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Field
	err := writer.WriteField("name", "Yadi Apriyadi")
	if err != nil {
		panic(err)
	}

	file, _ := writer.CreateFormFile("image", "ContohUpload.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
