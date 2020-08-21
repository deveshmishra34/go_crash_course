package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It is working")
}

func v1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is v1 route")
}

func serverHandler() {
	fmt.Printf("Server listening at 8000")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/v1", v1)
	fmt.Printf("Server listening at 8000")
	http.ListenAndServe(":8000", nil)
}
