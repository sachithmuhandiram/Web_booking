package main

import (
	"io"
	"net/http"
)

func dog(response http.ResponseWriter, req *http.Request) {
	io.WriteString(response, "Bow bow")
}

func cat(response http.ResponseWriter, req *http.Request) {
	io.WriteString(response, "Mew Mew")
}

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}
