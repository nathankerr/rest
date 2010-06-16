package main

import (
	"fmt"
	"http"
	"github.com/nathankerr/rest.go"
	"strconv"
)

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
