package request

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/NexFlare/build-http-server-go/internal/context"
	"github.com/NexFlare/build-http-server-go/internal/helper"
)

type Method string

const (
	GET Method = "GET"
	POST = "POST"
	PUT = "PUT"
	PATCH = "PATCH"
	DELETE = "DELETE"
)
type Request struct {
	RequestHeader
	URL string
	Method Method
	Body *string
	ctx *context.ServerContext
	conn net.Conn
}

type RequestHeader struct {
	Header map[string]string
}

func NewRequest(conn net.Conn) *Request{

	data, err := getData(conn)
	if err != nil {
		fmt.Println("error while reading request: ", err.Error())
		os.Exit(1)
	}
	headers, _ := getHeader(data)
	method, url, err :=getUrlAndMethod(data)
	var body *string
	if *method != string(GET) {
		body = getBody(data)
	}
	if err != nil {
		// TODO
		fmt.Println("Error while initializing request: ", err.Error())
	}


	return &Request{
		RequestHeader: *headers,
		URL: *url,
		Method: Method(*method),
		conn: conn,
		Body: body,
	}
}

func (r *Request) GetUrl() string {
	return r.URL
}

func (r *Request) GetMethod() Method {
	return r.Method
}

func (r *Request) GetContext() (*context.ServerContext) {
	return r.ctx
}

func (r *Request) GetConnection() net.Conn {
	return r.conn
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
		if len(headerSplit) > 1 {
			headers.Header[headerSplit[0]] = headerSplit[1][1:]
		}
	}
	return headers, nil
 }

 func getUrlAndMethod(data []byte) (*string, *string, error) {
	req := string(data[:])
	requestSplit := strings.Split(req, helper.CRLF)
	startLine := requestSplit[0]
	startSplit := strings.Split(startLine, " ")
	if len(startSplit) < 3 {
		return nil, nil, errors.New("invalid request")
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

func getBody(data []byte) (*string) {
	var body string
	req := string(data[:])
	requestSplit := strings.Split(req, helper.CRLF)
	if len(requestSplit) >= 3 {
		body = requestSplit[len(requestSplit)-1]
		body = strings.Trim(body, "\x00 ")
	}
	return &body
}