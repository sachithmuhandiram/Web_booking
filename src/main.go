package main

import(
	"fmt"
	"net/http"
)

type web_server int

func (w web_server) ServeHTTP(reswt http.ResponseWriter, req *http.Request){
	fmt.Fprintln(reswt ,"Simple server started")
}


func main(){

	var ws web_server

	http.ListenAndServe(":8080",ws)


}