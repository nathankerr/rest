package main

import (
	"fmt"
	"http"
	"log"
	"github.com/nathankerr/rest.go"
	"strconv"
)

type SnipsCollection struct {
	snips map[int]string
	nextId int
}

func NewSnipsCollection() (*SnipsCollection) { return &SnipsCollection{make(map[int]string), 0} }

func (snips *SnipsCollection) add(snip string) int {
	log.Stdoutf("SnipsCollection.add(%#v)", snip)
	var id = snips.nextId
	snips.snips[id] = snip
	snips.nextId = id + 1
	return id
}

func (snips *SnipsCollection) Index(c *http.Conn) {
	for k,v := range snips.snips {
		fmt.Fprintf(c, "<a href=\"%v\">%v</a>%v<br/>", k, k, v)
	}
}

func (snips *SnipsCollection) Find(c *http.Conn, idString string) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		rest.NotFound(c)
		return
	}

	snip, ok := snips.snips[id]
	if !ok {
		rest.NotFound(c)
		return
	}


	fmt.Fprintf(c, "<h1>Snip %v</h1><p>%v</p>", id, snip)
}

func main() {
	address := "127.0.0.1:3000"
	snips := NewSnipsCollection()

	snips.add("first post!")
	snips.add("me too")

	rest.Resource("snips", snips)

	http.ListenAndServe(address, nil)
}
