package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func HelloWithParamHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		fmt.Fprintf(writer, "Name: %s\n", name)
	} else {
		fmt.Fprintf(writer, "Hello World")
	}
}

func HelloWithFormDataHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firtName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	fmt.Fprintf(writer, "Hello %s %s\n", firtName, lastName)
}

// basic
// ini untuk contoh unit test tanpa menjalankan server
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	var response *http.Response = recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("BERHASIL DAPAT RESPONSE")
	bodyString := string(body)
	fmt.Println(bodyString)
}

// test with validasi param
func TestParamHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=ivans", nil)
	recorder := httptest.NewRecorder()

	HelloWithParamHandler(recorder, request)

	var response *http.Response = recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Response code is %d, expected %d", response.StatusCode, http.StatusOK)
	}

	if !strings.Contains(bodyString, "Name:") {
		t.Errorf("invalid response body, expected contains 'Name:'")
	}
	fmt.Println(bodyString)
}

func TestPostHttp(t *testing.T) {
	var requestBody = strings.NewReader("firstName=ivans&lastName=ardiansyah") // untuk format form data memang seperti ini, sama seperti get, hanya disimpan didalam body

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	HelloWithFormDataHandler(recorder, request)

	var response *http.Response = recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Response code is %d, expected %d", response.StatusCode, http.StatusOK)
	}

	fmt.Println(bodyString)
}
