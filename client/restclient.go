package main

import (
	"github.com/nathankerr/httplib.go"
	"log"
	"io/ioutil"
)

func main() {
	c := new(httplib.Client)
	response, err := c.Request("http://localhost:3000/snips/", "GET", nil, "")
	if err != nil {
		log.Exit(err)
	}

	log.Stdout(response)
	data,_ := ioutil.ReadAll(response.Body)
	log.Stdout(string(data))
}
