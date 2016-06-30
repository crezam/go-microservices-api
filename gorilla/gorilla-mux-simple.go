package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Camilo\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SimpleHandler)
	// Bind to a port and pass our router in
	http.ListenAndServe(":8445", r)
}
