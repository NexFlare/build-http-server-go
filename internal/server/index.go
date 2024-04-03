package server

import (
	"fmt"
	"net"
	"os"

	"github.com/NexFlare/build-http-server-go/internal/context"
	"github.com/NexFlare/build-http-server-go/internal/request"
)

type conn string

type server struct {
	ctx context.ServerContext
	listener net.Listener
}

func NewServer(l net.Listener) *server {
	return &server{
		listener: l,
	}
}

func CreateConnection(network string, address string) *server {

	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// defer l.Close()
	
	server := NewServer(l)
	return server
}

func(s *server) GetRequest(conn net.Conn) *request.Request {
	request := request.NewRequest(conn)
	return request
}

func(s *server) GetContext() context.ServerContext {
	return s.ctx
}

func(s *server) AcceptConnection() net.Conn{
	conn, err := s.listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}

func(s* server) StartServer(f func(conn net.Conn)) {
	for {
		conn := s.AcceptConnection()
		go f(conn)
	}
	
}