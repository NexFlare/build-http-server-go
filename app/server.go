package main

import (
	"fmt"
	"os"
	"strings"

	// Uncomment this block to pass the first stage

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	server := server.CreateConnection("tcp", "0.0.0.0:4221")
	request := server.GetRequest()
	url := request.GetUrl()
	ctx := server.GetContext()
	_conn, err := ctx.GetConnection()
	conn := (*_conn)
	if err != nil {
		os.Exit(1)
	}
	// path, err := GetUrl(conn)
	if url == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Contains(url, "/echo/") {
		content := (url)[6:]
		fmt.Println("Content is ", content)
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
	} else if url == "/user-agent"{
		content := request.Header["User-Agent"]
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	if err != nil {
		fmt.Println("Error while writing: ", err.Error())
		os.Exit(1)
	}
}