package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/NexFlare/build-http-server-go/internal/request"
	"github.com/NexFlare/build-http-server-go/internal/response"
	"github.com/NexFlare/build-http-server-go/internal/server"
)

func main() {

	fmt.Println("Logs from your program will appear here!")

	server := server.CreateConnection("tcp", "0.0.0.0:4221")
	dir := flag.String("directory", "", "The name of the directory")
	flag.Parse()
	server.StartServer(func(conn net.Conn) {
		request := server.GetRequest(conn)
		err := processRequest(request, dir)
		if err != nil {
			fmt.Println("Error while writing: ", err.Error())
			os.Exit(1)
		}
	})
	
}

func processRequest(req *request.Request, dir *string) error {
	url := req.GetUrl()
	conn := req.GetConnection()
	var err error
	res := response.NewResponse(conn)

	if url == "/" {
		res.Ok()
		return nil
	} else if strings.Contains(url, "/echo/") {
		content := (url)[6:]
		res.WriteHeader("Content-Type", "text/plain")
		res.SendWithBody(response.StatusOk, &content)
	} else if url == "/user-agent"{
		content := req.Header["User-Agent"]
		res.WriteHeader("Content-Type", "text/plain")
		res.SendWithBody(response.StatusOk, &content)
	} else if strings.Contains(url, "/files/"){
		fileName := url[7:]
		if dir != nil {
			file := fmt.Sprintf("%s%s", *dir, fileName)
			if req.Method == request.GET {
				data, readFileError := os.ReadFile(file)
				if readFileError != nil {
					res := response.NewResponse(conn)
					res.NotFound()
				} else {
					content := string(data[:])
					res.WriteHeader("Content-Type", "application/octet-stream")
					res.SendWithBody(response.StatusOk, &content)
				}
			} else if req.Method == request.POST {
				if req.Body == nil {
					res.BadRequest()
					return nil
				}
				f, err := os.Create(file)
				if err != nil {
					res.ServerError()
					return err
				}
				_, err = f.WriteString(*req.Body)
				if err != nil {
					res.ServerError()
					return err
				}
				res.Status = response.StatusCreated
				res.Send()
			}
			
		} else {
			res.NotFound()
		}

	}else {
		res.NotFound()
	}
	return err
}