package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.ListenAndServe(":8080", nil)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	fmt.Println(r.Method)
	fmt.Println(r.URL.String())
	fmt.Println(r.Header)
}
