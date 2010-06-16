package rest

import (
	"http"
	"net"
	"os"
)

type Client struct {
	conn *http.ClientConn
	resource *http.URL
}

// resource is the url to the base of the resource (i.e., http://127.0.0.1:3000/snips/)
func NewClient(resource string) (*Client, os.Error) {
	var client = new(Client)
	var err os.Error

	// setup host
	if client.resource, err = http.ParseURL(resource); err != nil {
		return nil, err
	}

	// Setup conn
	var tcpConn net.Conn
	if tcpConn, err = net.Dial("tcp", "", client.resource.Host); err != nil {
		return nil, err
	}
	client.conn = http.NewClientConn(tcpConn, nil)

	return client, nil
}

// GET /resource/
func (client *Client) Index() (*http.Response, os.Error) {
	var request *http.Request
	var response *http.Response
	var err os.Error

	request = new(http.Request)
	request.URL = client.resource
	request.Method = "GET"

	if err = client.conn.Write(request); err != nil {
		return nil, err
	}
	if response, err = client.conn.Read(); err != nil {
		return nil, err
	}

	return response, nil
}
