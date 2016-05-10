package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AcctHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(vars["id"]))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/account/{id}", AcctHandler)
	http.ListenAndServe(":8080", r)
}
