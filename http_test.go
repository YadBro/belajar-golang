package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	proto := response.Proto
	statusCode := response.StatusCode
	status := response.Status
	if err != nil {
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
	fmt.Println(proto)
	fmt.Println(status)
	fmt.Println(statusCode)
	require.Equal(t, bodyString, "Hello World")
	require.Equal(t, statusCode, 200)
}
