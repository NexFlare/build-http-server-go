package server

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/context"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
)

type conn string

type server struct {
	Request request.Request
	ctx context.ServerContext
}

func NewServer(ctx *context.ServerContext) *server {
	return &server{
		Request: *request.NewRequest(*ctx),
		ctx: *ctx,
	}
}

func CreateConnection(network string, address string) *server {

	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// ctx := context.TODO()
	defer l.Close()
	
	_conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	

	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}
	ctx := context.NewServerContext(_conn)
	server := NewServer(&ctx)
	return server
}

func(s *server) GetRequest() *request.Request {
	return &s.Request
}

func(s *server) GetContext() context.ServerContext {
	return s.ctx
}