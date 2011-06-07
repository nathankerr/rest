package rest

import (
	"bytes"
	"http"
	"io"
	"net"
	"os"
)

type Client struct {
	conn     *http.ClientConn
	resource *http.URL
}

// Creates a client for the specified resource.
//
// The resource is the url to the base of the resource (i.e.,
// http://127.0.0.1:3000/snips/)
func NewClient(resource string) (*Client, os.Error) {
	var client = new(Client)
	var err os.Error

	// setup host
	if client.resource, err = http.ParseURL(resource); err != nil {
		return nil, err
	}

	// Setup conn
	var tcpConn net.Conn
	if tcpConn, err = net.Dial("tcp", client.resource.Host); err != nil {
		return nil, err
	}
	client.conn = http.NewClientConn(tcpConn, nil)

	return client, nil
}

// Closes the clients connection
func (client *Client) Close() {
	client.conn.Close()
}

// General Request method used by the specialized request methods to create a request
func (client *Client) newRequest(method string, id string) (*http.Request, os.Error) {
	request := new(http.Request)
	var err os.Error

	request.ProtoMajor = 1
	request.ProtoMinor = 1
	request.TransferEncoding = []string{"chunked"}

	request.Method = method

	// Generate Resource-URI and parse it
	url := client.resource.String() + id
	if request.URL, err = http.ParseURL(url); err != nil {
		return nil, err
	}

	return request, nil
}

// Send a request
func (client *Client) Request(request *http.Request) (*http.Response, os.Error) {
	var err os.Error
	var response *http.Response

	// Send the request
	if err = client.conn.Write(request); err != nil {
		return nil, err
	}

	// Read the response
	if response, err = client.conn.Read(request); err != nil {
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

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() os.Error {
	return nil
}

// POST /resource
func (client *Client) Create(body string) (*http.Response, os.Error) {
	var request *http.Request
	var err os.Error

	if request, err = client.newRequest("POST", ""); err != nil {
		return nil, err
	}

	request.Body = nopCloser{bytes.NewBufferString(body)}

	return client.Request(request)
}

// PUT /resource/id
func (client *Client) Update(id string, body string) (*http.Response, os.Error) {
	var request *http.Request
	var err os.Error
	if request, err = client.newRequest("PUT", id); err != nil {
		return nil, err
	}

	request.Body = nopCloser{bytes.NewBufferString(body)}

	return client.Request(request)
}

// Parse a response-Location-URI to get the ID of the worked-on snip
func (client *Client) IdFromURL(urlString string) (string, os.Error) {
	var url *http.URL
	var err os.Error
	if url, err = http.ParseURL(urlString); err != nil {
		return "", err
	}

	return string(url.Path[len(client.resource.Path):]), nil
}

// DELETE /resource/id
func (client *Client) Delete(id string) (*http.Response, os.Error) {
	var request *http.Request
	var err os.Error
	if request, err = client.newRequest("DELETE", id); err != nil {
		return nil, err
	}

	return client.Request(request)
}
