package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"time"
	"strings"
	"io"
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

		go handler(conn)

	}

}

func handler(conn net.Conn){

	time_out := conn.SetDeadline(time.Now().Add(10* time.Second))

	if time_out != nil{
		log.Println("Timeout")
	}
	log.Println("Server started listening")

	// when a request comes pass to this
	request(conn)

	// after request parsed
	response(conn)

	defer conn.Close()
	
}

func request(conn net.Conn){
	/*
		This method will just parse METHOD and URL it passes
		When it receives an empty charactor, it breaks and go
		to response()

		i :=0, the request data is printed only on initial request
	*/
	i :=0

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		line := scanner.Text()

		if i ==0{
			method := strings.Fields(line)[0]
			url := strings.Fields(line)[1]

			fmt.Println("Method ->",method)
			fmt.Println("URL -> ",url)
			// later here will update for handling multiple urls
			if url == "/hello"{
				hello_response(conn)
			}
		}
		if line == ""{
			break		// empty charactor
		}
		i++

	}
} 
// end of request function

func response(conn net.Conn){

	fmt.Fprintln(conn,"This is response")

	defer conn.Close() 
}

// first route
func hello_response(conn net.Conn){
	io.WriteString(conn,"Hello from response\n")

	defer conn.Close() 
}