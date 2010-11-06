// Wraps the http package with a HTTP method and header aware muxer.
// Code derived from the http package implementation of DefaultServeMux.

package rest

import (
	"fmt"
	"http"
	"strings"
)

var resources = make (map[string]interface{})

// Lists all the items in the resource
// GET /resource/
type Index interface {
	Index(http.ResponseWriter)
}

// Creates a new resource item
// POST /resource/
type Create interface {
	Create(http.ResponseWriter, *http.Request)
}

// Views a resource item
// GET /resource/id
type Find interface {
	Find(http.ResponseWriter, string)
}

// PUT /resource/id
type Update interface {
	Update(http.ResponseWriter, string, *http.Request)
}

// DELETE /resource/id
type Delete interface {
	Delete(http.ResponseWriter, string)
}

// Return options to use the service. If string is nil, then it is the base URL
// OPTIONS /resource/id
// OPTIONS /resource/
type Options interface {
	Options(http.ResponseWriter, string)
}

func resourceHandler(c http.ResponseWriter, req *http.Request) {
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
			if resIndex, ok := resource.(Index); ok {
				resIndex.Index(c)
			} else {
				NotImplemented(c)
			}
		case "POST":
			// Create
			if resCreate, ok := resource.(Create); ok {
				resCreate.Create(c, req)
			} else {
				NotImplemented(c)
			}
		case "OPTIONS":
			// automatic options listing
			if resOptions, ok := resource.(Options); ok {
				resOptions.Options(c, id)
			} else {
				NotImplemented(c)
			}
		default:
			NotImplemented(c)
		}
	} else {
		switch req.Method {
		case "GET":
			// Find
			if resFind, ok := resource.(Find); ok {
				resFind.Find(c, id)
			} else {
				NotImplemented(c)
			}
		case "PUT":
			// Update
			if resUpdate, ok := resource.(Update); ok {
				resUpdate.Update(c, id, req)
			} else {
				NotImplemented(c)
			}
		case "DELETE":
			// Delete
			if resDelete, ok := resource.(Delete); ok {
				resDelete.Delete(c, id)
			} else {
				NotImplemented(c)
			}
		case "OPTIONS":
			// automatic options
			if resOptions, ok := resource.(Options); ok {
				resOptions.Options(c, id)
			} else {
				NotImplemented(c)
			}
		default:
			NotImplemented(c)
		}
	}
}

func Resource(name string, res interface{}) {
	resources[name] = res
	http.Handle("/" + name + "/", http.HandlerFunc(resourceHandler))
}

func NotFound(c http.ResponseWriter) {
	http.Error(c, "404 Not Found", http.StatusNotFound)
}

func NotImplemented(c http.ResponseWriter) {
	http.Error(c, "501 Not Implemented", http.StatusNotImplemented)
}

func Created(c http.ResponseWriter, location string) {
	c.SetHeader("Location", location)
	http.Error(c, "201 Created", http.StatusCreated)
}

func Updated(c http.ResponseWriter, location string) {
	c.SetHeader("Location", location)
	http.Error(c, "200 OK", http.StatusOK)
}

func BadRequest(c http.ResponseWriter, instructions string) {
	c.WriteHeader(http.StatusBadRequest)
	c.Write([]byte(instructions))
}

func NoContent(c http.ResponseWriter) {
	http.Error(c, "204 No Content", http.StatusNoContent)
}
