package tcp_client

import (
	"bufio"
	"net"
	"strings"
)

type Client struct {
	conn         net.Conn
	serverReader *bufio.Reader
}

func NewClient(addr string) (client *Client, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	return &Client{
		conn:         conn,
		serverReader: bufio.NewReader(conn),
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SendMessage(message string) error {
	_, err := c.conn.Write([]byte(strings.TrimSpace(message) + "\n"))
	return err
}

func (c *Client) GetMessage() (string, error) {
	response, err := c.serverReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(response), nil
}
