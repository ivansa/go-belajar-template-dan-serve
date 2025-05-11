package belajar_golang_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandlerSetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Name-Server"
	cookie.Value = "Ivans Laptop"
	cookie.Path = "/"

	http.SetCookie(writer, cookie)

	fmt.Fprintf(writer, "INI COOKIE SERVER %s\n", cookie.String())

}

func HandlerGetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Name-Server")
	if err != nil {
		fmt.Fprintf(writer, "COOKIE NOT FOUND\n")
	} else {
		fmt.Fprintf(writer, "INI COOKIE CLIENT %s\n", cookie.String())
	}
}

func TestCoockie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", HandlerSetCookie)
	mux.HandleFunc("/get", HandlerGetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestCookieSet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set", nil)
	recorder := httptest.NewRecorder()

	HandlerSetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("INI COOKIE %s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestCookieGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get", nil)
	request.AddCookie(&http.Cookie{Name: "X-Name-Server", Value: "Ivans Laptop"})
	recorder := httptest.NewRecorder()
	HandlerGetCookie(recorder, request)

	body := recorder.Body.String()
	fmt.Println(body)
}
