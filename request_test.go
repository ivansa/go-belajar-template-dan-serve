package belajar_golang_web

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// http://localhost:8080/?name=ivans&name=ardiansyah
func TestRequestWithParam(t *testing.T) {
	var handler http.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Request Method: %s\n", request.Method)
		fmt.Fprintf(writer, "Request URI: %s\n", request.RequestURI)

		// get parameter name pertama
		name := request.URL.Query().Get("name")

		// get all parameter
		names := request.URL.Query()["name"]

		if name != "" {
			fmt.Fprintf(writer, "Name: %s\n", name)
			fmt.Fprintf(writer, "FullName: %s\n", strings.Join(names, " "))

		} else {
			fmt.Fprintf(writer, "Name not found\n")
		}
	})

	startServer(handler)
}

func TestPostRequest(t *testing.T) {
	var handler http.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Request Method: %s\n", request.Method)
		fmt.Fprintf(writer, "Request URI: %s\n", request.RequestURI)

		if request.Method != "POST" {
			panic("ERROR: Method not allowed")
		}

		// x-www-form-urlencoded
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}

		firstName := request.PostForm.Get("first_name")
		lastName := request.PostForm.Get("last_name")

		fmt.Fprintf(writer, "Name: %s %s \n", firstName, lastName)
	})

	startServer(handler)
}

func TestRequest(t *testing.T) {
	var handler http.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("X-Powered-By", "Ivans")

		fmt.Fprintf(writer, "Request Method: %s\n", request.Method)
		fmt.Fprintf(writer, "Request URI: %s\n", request.RequestURI)
		fmt.Fprintf(writer, "Content-Type: %s\n", request.Header.Get("content-type"))
		fmt.Fprintf(writer, "accept: %s\n", request.Header.Get("accept"))
	})

	startServer(handler)
}

func startServer(handler http.Handler) {
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
