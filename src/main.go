package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	listen,err := net.Listen("tcp",":8080")

	if err != nil{
		log.Panic(err)
	}

	defer listen.Close()

	for {
		conn,err := listen.Accept()

		if err != nil {
			log.Println(err)
		}
		log.Println("Server started")
		io.WriteString(conn, "\n Hello\n")
		fmt.Fprintln(conn,"How are you?")
		conn.Close()
	}
}