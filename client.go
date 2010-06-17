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

func (client *Client) newRequest(method string, id string) (*http.Request, os.Error) {
	request := new(http.Request)
	var err os.Error

	url := client.resource.String() + id
	if request.URL, err = http.ParseURL(url); err != nil {
		return nil, err
	}

	return request, nil
}

func (client *Client) Request(request *http.Request) (*http.Response, os.Error) {
	var err os.Error
	var response *http.Response

	if err = client.conn.Write(request); err != nil {
		return nil, err
	}

	if response, err = client.conn.Read(); err != nil {
		return nil, err
	}

	return response, nil
}

// GET /resource/
func (client *Client) Index() (*http.Response, os.Error) {
	var request *http.Request
	var err os.Error

	if request, err = client.newRequest("GET", ""); err != nil {
		return nil, err
	}

	return client.Request(request)
}

// GET /resource/id
func (client *Client) Find(id string) (*http.Response, os.Error) {
	var request *http.Request
	var err os.Error

	if request, err = client.newRequest("GET", id); err != nil {
		return nil, err
	}

	return client.Request(request)
}
