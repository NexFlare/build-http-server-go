package response

import (
	"fmt"
	"net"
	"strings"
)

type Response struct {
	Header map[string]string
	Protocol string
	Status string
	Body *string
	Conn net.Conn
}

func NewResponse(conn net.Conn) Response {
	return Response{
		Conn: conn,
		Protocol: "HTTP/1.1",
	}
}

func (r *Response) Send() {
	var sb strings.Builder
	defer r.Conn.Close()
	if r.Protocol == "" {
		r.Protocol = "HTTP/1.1"
	}
	str := fmt.Sprintf("%v %v\r\n", r.Protocol, r.Status)
	sb.WriteString(str)
	if r.Header != nil {
		for k,v := range r.Header {
			str := fmt.Sprintf("%s:%s\r\n", k,v)
			sb.WriteString(str)
		}
		sb.WriteString("\r\n")
	}
	if r.Body != nil {
		sb.WriteString(*r.Body)
		sb.WriteString("\r\n")
	}
	sb.WriteString("\r\n")
	finalBody := sb.String()
	fmt.Println("Final body is", finalBody)
	r.Conn.Write([]byte(finalBody))
}

func (r *Response) NotFound() {
	r.Status = "404 Not Found"
	r.Send()
}

func (r *Response) Ok() {
	r.Status = "200 OK"
	r.Send()
}

func (r *Response) ServerError() {
	r.Status = "500 Internal Server Error"
	r.Send()
}

func (r *Response) BadRequest() {
	r.Status = "400 Bad Request"
	r.Send()
}