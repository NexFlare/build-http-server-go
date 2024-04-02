package context

import (
	"context"
	"errors"
	"net"
)


type ServerContext struct {
	context.Context
}

func NewServerContext(conn net.Conn) ServerContext {
	ctx := context.Background()

	return ServerContext{
		context.WithValue(ctx, "conn", &conn),
	}
	
}

func(s *ServerContext) GetConnection() (*net.Conn, error) {
	conn, ok := s.Value("conn").(*net.Conn)
	if !ok {
		return nil, errors.New("unable to get conn")
	}
	return conn, nil
}