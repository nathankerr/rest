/*
	Defines SnipsCollection struct-type and NewSnipsCollection()
	It is used by the example REST server for storing and managing snips
	It has the methods Add, WithId, All and Remove.
*/
package main

import (
	"log"
)

// Snip definintion
type Snip struct {
	Id   int
	Body string
}

func NewSnip(id int, body string) *Snip {
	log.Println("Creating new Snip:", id, body)
	return &Snip{id, body}
}

// SnipsCollection definition
type SnipsCollection struct {
	snips  map[int]string
	nextId int
}

func NewSnipsCollection() *SnipsCollection {
	log.Println("Creating new SnipsCollection")
	return &SnipsCollection{make(map[int]string), 0}
}

func (c *SnipsCollection) Add(body string) int {
	log.Println("Adding Snip:", body)
	id := c.nextId
	c.nextId++

	c.snips[id] = body

	return id
}

func (c *SnipsCollection) WithId(id int) (*Snip, bool) {
	log.Println("Finding Snip with id: ", id)
	body, ok := c.snips[id]
	if !ok {
		return nil, false
	}

	return &Snip{id, body}, true
}

func (c *SnipsCollection) All() []*Snip {
	log.Println("Finding all Snips")
	all := make([]*Snip, len(c.snips))

	for id, body := range c.snips {
		all[id] = &Snip{id, body}
	}

	return all
}

func (c *SnipsCollection) Remove(id int) {
	delete(c.snips, id)
}
