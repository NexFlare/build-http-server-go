package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

type void struct{}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()
	
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
	}
	req := string(buffer[:])
	reqLines := strings.Split(req, "\r\n")
	startLine := reqLines[0]
	path := strings.Split(startLine, " ")[1]
	set := make(map[string]void)
	set["/"] = void{}
	if _, ok:= set[path]; ok {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	if err != nil {
		fmt.Println("Error while writing: ", err.Error())
		os.Exit(1)
	}
}

// func
