package rest

import (
	"http"
	"net"
	"os"
)

type Client struct {
	conn *http.ClientConn
	host string
}

func NewClient(host http.URL) (*Client, os.Error) {
	var client = new(Client)

	// Setup conn
	var tcpConn net.Conn
	var err os.Error
	if tcpConn, err = net.Dial("tcp", "", client.host); err != nil {
		return nil, err
	}
	client.conn = http.NewClientConn(tcpConn, nil)

	// setup host
	client.host = host.Host

	return client, nil
}
