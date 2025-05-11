package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// Ini adalah contoh file static jika perlu image css js dll
func TestFileServer(t *testing.T) {
	dir := http.Dir("./assets")
	fileServer := http.FileServer(dir)

	handler := http.StripPrefix("/static/", fileServer) // jika ingin menghilangkan prefix

	mux := http.NewServeMux()
	mux.Handle("/static/", handler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed assets
var assets embed.FS

func TestFileServerEmbed(t *testing.T) {
	subFS, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}
	dir := http.FS(subFS)
	fileServer := http.FileServer(dir)

	handler := http.StripPrefix("/static/", fileServer) // jika ingin menghilangkan prefix

	mux := http.NewServeMux()
	mux.Handle("/static/", handler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
