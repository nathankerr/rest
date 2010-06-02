package main

import (
	"fmt"
	"http"
	"log"
	"rest"
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
	log.Stdout("SnipsCollection.Index(%#v)", c)

	for k,v := range snips.snips {
		fmt.Fprintf(c, "<a href=\"%v\">%v</a>%v<br/>", k, k, v)
	}
}

func (snips *SnipsCollection) Find(c *http.Conn, idString string) {
	log.Stdout("SnipsCollection.Find(%#v, %#v)", c, idString)
	var id, _ = strconv.Atoi(idString)
	var snip, ok = snips.snips[id]
	if ok {
		fmt.Fprintf(c, "<h1>Snip %v</h1><p>%v</p>", id, snip)
	} else {
		rest.NotFound(c)
		return
	}
}

type OtherCollection struct {
	name string
}

func main() {
	var snips = &SnipsCollection{make(map[int]string), 0}
	var other = &OtherCollection{"other"}

	snips.add("first post!")
	snips.add("me too")

	rest.Resource("snips", snips)
	rest.Resource("other", other)

	http.ListenAndServe("127.0.0.1:3000", nil)
}
