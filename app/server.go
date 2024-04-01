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
	path, err := GetUrl(conn)
	// set := make(map[string]void)
	// set["/"] = void{}
	if *path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Contains(*path, "/echo/") {
		content := (*path)[6:]
		fmt.Println("Content is ", content)
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	if err != nil {
		fmt.Println("Error while writing: ", err.Error())
		os.Exit(1)
	}
}

// func

func GetUrl(conn net.Conn) (*string, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return nil, err
	}
	req := string(buffer[:])
	reqLines := strings.Split(req, "\r\n")
	startLine := reqLines[0]
	path := strings.Split(startLine, " ")[1]
	return &path, nil
}