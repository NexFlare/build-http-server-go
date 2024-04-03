package response

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Response struct {
	Header map[string]string
	Protocol string
	Status Status
	Body *string
	Conn net.Conn
}

func NewResponse(conn net.Conn) Response {
	return Response{
		Conn: conn,
		Protocol: "HTTP/1.1",
		Header: map[string]string{},
	}
}

func (r *Response) SendWithBody(status Status, body *string) {
	r.Body = body
	r.Status = status
	fmt.Printf("Respose obj is %+v", r)
	r.Send()
}

func (r *Response) Send() {
	var sb strings.Builder
	defer r.Conn.Close()
	if r.Protocol == "" {
		r.Protocol = "HTTP/1.1"
	}
	if r.Body != nil {
		r.Header["Content-Length"] = strconv.Itoa(len(*r.Body))
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
	r.Conn.Write([]byte(finalBody))
}

func (r *Response) WriteHeader(key string, value string) {
	r.Header[key] = value
}

func (r *Response) NotFound() {
	r.Status = StatusNotFound
	r.Send()
}

func (r *Response) Ok() {
	r.Status = StatusOk
	r.Send()
}

func (r *Response) ServerError() {
	r.Status = StatusInternalServerError
	r.Send()
}

func (r *Response) BadRequest() {
	r.Status = StatusBadRequest
	r.Send()
}