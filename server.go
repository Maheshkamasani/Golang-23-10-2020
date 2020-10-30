package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/hello", h2)
	http.ListenAndServe(":8080", nil)
}

func h(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}