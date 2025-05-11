package belajar_golang_web

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCodeHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Name not found"))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello " + name))
	}
}

func TestResponseCodeServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ResponseCodeHandler),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=ivans", nil)
	recorder := httptest.NewRecorder()

	ResponseCodeHandler(recorder, request)

	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)
	if bodyString != "Hello ivans" {
		t.Errorf("Expected body 'Hello ivans', got '%s'", bodyString)
	}
}
