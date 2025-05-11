package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini Home")
	})

	mux.HandleFunc("/hi/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi anonymous")
	})

	mux.HandleFunc("/hi/ivans/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi ivans")
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
