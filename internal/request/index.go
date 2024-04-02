package request

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/context"
	"github.com/codecrafters-io/http-server-starter-go/internal/helper"
)

type Request struct {
	RequestHeader
	URL string
	Method string
}

type RequestHeader struct {
	Header map[string]string
}

func NewRequest(ctx context.ServerContext) *Request{

	conn, err := ctx.GetConnection()
	if err != nil {
		fmt.Println("Error while getting connection: ", err.Error())
	}
	data, err := getData(*conn)
	if err != nil {
		fmt.Println("error while reading request: ", err.Error())
		os.Exit(1)
	}
	headers, _ := getHeader(data)
	method, url, err :=getUrlAndMethod(data)
	if err != nil {
		// TODO
		fmt.Println("Error while initializing request: ", err.Error())
	}

	return &Request{
		RequestHeader: *headers,
		URL: *url,
		Method: *method,
	}
}

func (r *Request) GetUrl() string {
	return r.URL
}

func (r *Request) GetMethod() string {
	return r.Method
}


func getHeader(data []byte) (*RequestHeader, error) {
	req := string(data[:])
	requestSplit := strings.Split(req, helper.CRLF)
	headers :=  &RequestHeader{
		Header: map[string]string{},
	}
	if len(requestSplit) == 3 {
		return headers, nil
	}

	requestSplit = requestSplit[1:len(requestSplit)-2]

	for i:=0;i<len(requestSplit);i++ {
		header := requestSplit[i]
		headerSplit := strings.Split(header, ":")
		headers.Header[headerSplit[0]] = headerSplit[1][1:];
	}
	return headers, nil
 }

 func getUrlAndMethod(data []byte) (*string, *string, error) {
	req := string(data[:])
	requestSplit := strings.Split(req, helper.CRLF)
	startLine := requestSplit[0]
	startSplit := strings.Split(startLine, " ")
	if len(startSplit) < 3 {
		return nil, nil, errors.New("Invalid request")
	}
	return &startSplit[0], &startSplit[1], nil
 }

 func getData(conn net.Conn) ([]byte, error ){
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}