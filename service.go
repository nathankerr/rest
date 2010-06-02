package main

import (
	"fmt"
	"http"
	"log"
	"rest"
)

var Snips = make(map[int]string)
var NextId = 0

func AddSnip(snip string) int {
	var id = NextId
	Snips[id] = snip
	NextId = NextId + 1
	return id
}

func SnipsList(c *http.Conn, req *http.Request) {
	log.Stdout("SnipsList")
	c.SetHeader("Content-Type", "text")
	for k,v := range Snips {
		fmt.Fprintf(c, "%v: %v\n", k, v)
	}
}

func Snip(c *http.Conn, req *http.Request) {
	
}

func main() {
	AddSnip("test 1")
	AddSnip("2")

	rest.Handle("/snip/", http.HandlerFunc(SnipsList), []string{"GET"}, nil)
	rest.Handle("/snip/(.*)", http.HandlerFunc(CreateSnip), []string{"POST"}, nil)

	rest.ListenAndServe("127.0.0.1:3000", nil)
}
