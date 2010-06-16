package main

import (
	"log"
)

type SnipsCollection struct {
	snips map[int]string
	nextId int
}

func NewSnipsCollection() (*SnipsCollection) { return &SnipsCollection{make(map[int]string), 0} }

func (snips *SnipsCollection) Add(snip string) int {
	log.Stdoutf("SnipsCollection.add(%#v)", snip)
	var id = snips.nextId
	snips.snips[id] = snip
	snips.nextId = id + 1
	return id
}
