package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// menggunakan ini hanya bisa menggunakan 1 handler saja / url
func TestHandler(t *testing.T) {
	var handler http.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
