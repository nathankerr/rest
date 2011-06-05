/*
	Defines SnipsCollection struct-type and NewSnipsCollection()
	SnipsCollection has methods Add, WithId, All and Remove.
*/
package main

import (
	"log"
	"container/vector"
)

// Snip definintion
type Snip struct {
	Id int
	Body string
}

func NewSnip(id int, body string) *Snip {
	log.Println("Creating new Snip:", id, body)
	return &Snip{id, body}
}

// SnipsCollection definition
type SnipsCollection struct {
	v *vector.Vector
	nextId int
}

func NewSnipsCollection() (*SnipsCollection) {
	log.Println("Creating new SnipsCollection")
	return &SnipsCollection{new(vector.Vector), 0}
}

func (snips *SnipsCollection) Add(body string) int {
	log.Println("Adding Snip:", body)
	id := snips.nextId
	snips.nextId++

	snip := NewSnip(id, body)
	snips.v.Push(snip)

	return id
}

func (snips *SnipsCollection) WithId(id int) (*Snip, bool) {
	log.Println("Finding Snip with id: ", id)
	all := *snips.v
	for _, v := range all {
		snip, ok := v.(*Snip)
		if ok && snip.Id == id {
			return snip, true
		}
	}
	return nil, false
}

func (snips *SnipsCollection) All() []*Snip {
	log.Println("Finding all Snips")
	data := *snips.v
	all := make([]*Snip, len(data))

	for k, v := range data {
		if snip, ok := v.(*Snip); ok {
			all[k] = snip
		}
	}

	return all
}

func (snips *SnipsCollection) Remove(id int) {
	length := snips.v.Len()
	for i := 0; i < length; i++ {
		snip := snips.v.At(i).(*Snip)
		if snip.Id == id {
			snips.v.Delete(i)
			return
		}
	}
}
