package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage

	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
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
			defer conn.Close()
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
	url := request.GetUrl()
	conn := request.GetConnection()
	var err error
	if url == "/" {
		res := response.NewResponse(conn)
		res.Ok()
	} else if strings.Contains(url, "/echo/") {
		content := (url)[6:]
		res := response.Response{
			Header: map[string]string{
				"Content-Type": "text/plain",
				"Content-Length": strconv.Itoa(len(content)),
			},
			Conn: conn,
			Body: &content,
		}
		res.Ok()
	} else if url == "/user-agent"{
		content := request.Header["User-Agent"]
		res := response.Response{
			Header: map[string]string{
				"Content-Type": "text/plain",
				"Content-Length": strconv.Itoa(len(content)),
			},
			Conn: conn,
			Body: &content,
		}
		res.Ok()
	} else if strings.Contains(url, "/files/"){
		fileName := url[7:]
		if dir != nil {
			file := fmt.Sprintf("%s%s", *dir, fileName)
			if request.Method == "GET" {
				data, readFileError := os.ReadFile(file)
				if readFileError != nil {
					res := response.NewResponse(conn)
					res.NotFound()
				} else {
					content := string(data[:])
					res := response.Response{
						Header: map[string]string{
							"Content-Type": "application/octet-stream",
							"Content-Length": strconv.Itoa(len(content)),
						},
						Conn: conn,
						Body: &content,
					}
					res.Ok()
				}
			} else if request.Method == "POST" {
				res := response.NewResponse(conn)
				if request.Body == nil {
					res.BadRequest()
					return nil
				}
				f, err := os.Create(file)
				if err != nil {
					res.ServerError()
					return err
				}
				_, err = f.WriteString(*request.Body)
				if err != nil {
					res.ServerError()
					return err
				}
				res.Status = "201 Created"
				res.Send()
			}
			
		} else {
			res := response.NewResponse(conn)
			res.NotFound()
		}

	}else {
		res := response.NewResponse(conn)
		res.NotFound()
	}
	return err
}