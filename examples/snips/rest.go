/*
	Defines Index, Find, Create, Update and Delete methods for SnipsCollection
	These will be called on the corresponding REST requests
*/
package main

import (
	"fmt"
	"http"
	"io/ioutil"
	"os"
	"github.com/Swoogan/rest.go"
	"strconv"
)

// String used as error message on invalid formatting (request body could not be parsed into a snip)
var formatting = "formatting instructions go here"

// Get an index of the snips in the collection
func (snips *SnipsCollection) Index(c http.ResponseWriter) {
	for _,snip := range snips.All() {
		fmt.Fprintf(c, "<a href=\"%v\">%v</a>%v<br/>", snip.Id, snip.Id, snip.Body)
	}
}

// Find a snip from the collection, identified by the ID
func (snips *SnipsCollection) Find(c http.ResponseWriter, idString string) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		rest.NotFound(c)
		return
	}

	snip, ok := snips.WithId(id)
	if !ok {
		rest.NotFound(c)
		return
	}

	fmt.Fprintf(c, "<h1>Snip %v</h1><p>%v</p>", snip.Id, snip.Body)
}

// Create and add a new snip to the collection
func (snips *SnipsCollection) Create(c http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		rest.BadRequest(c, formatting)
		return
	}

	id := snips.Add(string(data))
	rest.Created(c, fmt.Sprintf("%v%v", request.URL.String(), id))
}

// Update a snip identified by an ID with the data sent as request-body
func (snips *SnipsCollection) Update(c http.ResponseWriter, idString string, request *http.Request) {
	// Parse ID of type string to int
	var id int
	var err os.Error
	if id, err = strconv.Atoi(idString); err != nil {
		// The ID could not be converted from string to int
		rest.NotFound(c)
		return
	}

	// Find the snip with the ID
	var snip *Snip
	var ok bool
	if snip, ok = snips.WithId(id); !ok {
		// A snip with the passed ID could not be found in our collection
		rest.NotFound(c)
	}

	// Get the request-body for data to update the snipped to
	var data []byte
	if data, err = ioutil.ReadAll(request.Body); err != nil {
		// The request body could not be read, thus it was a bad request
		rest.BadRequest(c, formatting)
		return
	}

	// Set the snips body
	snip.Body = string(data)
	// Respond to indicate successful update
	rest.Updated(c, request.URL.String())
}

// Delete a snip identified by ID from the collection
func (snips *SnipsCollection) Delete(c http.ResponseWriter, idString string) {
	var id int
	var err os.Error
	if id, err = strconv.Atoi(idString); err != nil {
		rest.NotFound(c)
	}

	snips.Remove(id)
	rest.NoContent(c)
}
