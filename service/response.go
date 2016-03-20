package service

import (
	"net"
)

type Response struct {
	conn  net.Conn
}

func NewResponse(cx net.Conn) *Response {
	return &Response{
		conn: cx,
	}
}

func (resp Response) Send(message string) (int, error) {
	return resp.conn.Write([]byte(message + "\n"))
}
