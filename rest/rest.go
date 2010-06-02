// Wraps the http package with a HTTP method and header aware muxer.
// Code derived from the http package implementation of DefaultServeMux.

package rest

import (
	"fmt"
	"http"
	"log"
	"strings"
)

var resources = make (map[string]interface{})

// Lists all the items in the resource
type Index interface {
	Index(*http.Conn)
}

// Creates a new resource item
type Create interface {
	Create(*http.Conn, map[string]string)
}

// Views a resource item
type Find interface {
	Find(*http.Conn, string)
}

type Update interface {
	Update(*http.Conn, string, map[string]string)
}

type Delete interface {
	Delete(*http.Conn, string)
}

func resourceHandler(c *http.Conn, req *http.Request) {
	log.Stdoutf("rest.resourceHandler(%#v, %#v)", c, req)

	var resourceEnd = strings.Index(req.URL.Path[1:], "/") + 1
	var resourceName string
	if (resourceEnd == -1) {
		resourceName = req.URL.Path[1:]
	} else {
		resourceName = req.URL.Path[1:resourceEnd]
	}
	var id = req.URL.Path[resourceEnd+1:]

	resource, ok := resources[resourceName]
	if !ok {
		fmt.Fprintf(c, "resource %s not found\n", resourceName)
	}

	if len(id) == 0 {
		switch req.Method {
		case "GET":
			// Index
			var resIndex, ok = resource.(Index)
			if ok {
				resIndex.Index(c)
			} else {
				NotImplemented(c)
			}
		case "POST":
			// Create
			NotImplemented(c)
		case "OPTIONS":
			// automatic options listing
			NotImplemented(c)
		default:
			NotImplemented(c)
		}
	} else {
		switch req.Method {
		case "GET":
			// Find
			var resFind, ok = resource.(Find)
			if ok {
				resFind.Find(c, id)
			} else {
				NotImplemented(c)
			}
		case "PUT":
			// Update
			NotImplemented(c)
		case "DELETE":
			// Delete
			NotImplemented(c)
		case "OPTIONS":
			// automatic options
			NotImplemented(c)
		default:
			NotImplemented(c)
		}
	}
}

func Resource(name string, res interface{}) {
	log.Stdoutf("rest.Resource(%#v, %#v)", name, res)

	resources[name] = res
	http.Handle("/" + name + "/", http.HandlerFunc(resourceHandler))
}

func NotFound(c *http.Conn) {
	http.Error(c, "404 Not Found", 404)
}

func NotImplemented(c *http.Conn) {
	http.Error(c, "501 Not Implemented", 501)
}
