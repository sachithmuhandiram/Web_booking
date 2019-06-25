package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)                           // handle URL, for this use foo func
	http.Handle("/favicon.ico", http.NotFoundHandler()) // icon on top of tab
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	query := req.FormValue("q")                   // URL should look like http://localhost:8080/?q=hello
	io.WriteString(w, "You serched for : "+query) // query = hello
}
