package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SimpleHandler)
	// Bind to a port and pass our router in
	http.ListenAndServe(":8080", r)
}