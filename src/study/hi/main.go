package main

import (
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)

	log.Println("Starting httpserver")
	log.Fatal(http.ListenAndServe(":18080", mux))
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, " + r.RemoteAddr))
}
