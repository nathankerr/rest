package main

import (
	"fmt"
	"http"
	"io/ioutil"
	"os"
	"github.com/nathankerr/rest.go"
	"strconv"
)

var formatting = "formatting instructions go here"

func (snips *SnipsCollection) Index(c *http.Conn) {
	for _,snip := range snips.All() {
		fmt.Fprintf(c, "<a href=\"%v\">%v</a>%v<br/>", snip.Id, snip.Id, snip.Body)
	}
}

func (snips *SnipsCollection) Find(c *http.Conn, idString string) {
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

func (snips *SnipsCollection) Create(c *http.Conn, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		rest.BadRequest(c, formatting)
		return
	}

	id := snips.Add(string(data))
	rest.Created(c, fmt.Sprintf("%v%v", request.URL.String(), id))
}

func (snips *SnipsCollection) Update(c *http.Conn, idString string, request *http.Request) {
	var id int
	var err os.Error
	if id, err = strconv.Atoi(idString); err != nil {
		rest.NotFound(c)
		return
	}

	var snip *Snip
	var ok bool
	if snip, ok = snips.WithId(id); !ok {
		rest.NotFound(c)
	}

	var data []byte
	if data, err = ioutil.ReadAll(request.Body); err != nil {
		rest.BadRequest(c, formatting)
		return
	}

	snip.Body = string(data)
	rest.Updated(c, request.URL.String())
}
