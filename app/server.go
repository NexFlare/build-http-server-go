package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	// Uncomment this block to pass the first stage

	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	server := server.CreateConnection("tcp", "0.0.0.0:4221")
	dir := flag.String("directory", "", "The name of the directory")
	flag.Parse()
	for {
		conn := server.AcceptConnection()
		go func() {
			request := server.GetRequest(conn)
			err := processRequest(request, dir)
			if err != nil {
				fmt.Println("Error while writing: ", err.Error())
				os.Exit(1)
			}
		}()
	}
}

func processRequest(request *request.Request, dir *string) error {
	if dir != nil {
		fmt.Println("The directory name is", *dir)
	}
	
	url := request.GetUrl()
	conn := request.GetConnection()
	var err error
	if url == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Contains(url, "/echo/") {
		content := (url)[6:]
		fmt.Println("Content is ", content)
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
	} else if url == "/user-agent"{
		content := request.Header["User-Agent"]
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
	} else if strings.Contains(url, "/files/"){
		fileName := url[7:]
		fmt.Println("file name is", fileName)
		if dir != nil {
			file := fmt.Sprintf("%s%s", *dir, fileName)
			data, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("Error is", err.Error())
				_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			} else {
				fmt.Println("File size is", len(data))
				content := string(data[:])
				_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length:%+v\r\n\r\n%v\r\n",len(content), content)))
			}
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		}

	}else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
	return err
}