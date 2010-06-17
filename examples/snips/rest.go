package main

import (
	"fmt"
	"http"
	"github.com/nathankerr/rest.go"
	"strconv"
)

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
