package main

import (
	"embed"
	"log"
	"net/http"
)

var (
	HTTP_PORT = "0.0.0.0:8080"
)

//go:embed index.html
var contentfs embed.FS

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.FS(contentfs))
	mux.Handle("/", http.StripPrefix("/", fs))
	mux.HandleFunc("/ping", Ping)

	log.Println("Run Server:", HTTP_PORT)
	http.ListenAndServe(HTTP_PORT, mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{"msg":"pong"}`))
}
