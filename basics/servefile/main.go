package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/rosi.jpg", rosiPic) // Rosi is my dog :-)
	http.ListenAndServe(":8080", nil)
}

func dog(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, `<img src="rosi.jpg">`)
}

func rosiPic(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "rosi.jpg")
}
