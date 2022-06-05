package tcp_server

import (
	"net"
)

type Handler interface {
	ServeTCP(conn net.Conn)
}

// ListenAndServe incoming TCP connection
func ListenAndServe(addr string, handler Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handler.ServeTCP(conn)
	}
}
