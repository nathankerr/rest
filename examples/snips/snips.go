package main

import (
	"flag"
	"fmt"
	"http"
	"io/ioutil"
	"log"
	"os"
	"github.com/nathankerr/rest.go"
)

var server = flag.Bool("server", false, "start in server mode")

func main() {
	flag.Parse()

	if *server {
		serve()
	} else {
		client()
	}
}

func serve() {
	log.Stdout("Starting Server")
	address := "127.0.0.1:3000"
	snips := NewSnipsCollection()

	snips.Add("first post!")
	snips.Add("me too")

	rest.Resource("snips", snips)

	http.ListenAndServe(address, nil)
}

func client() {
	log.Stdout("Starting Client")
	var snips *rest.Client
	var err os.Error

	if snips, err = rest.NewClient("http://127.0.0.1:3000/snips/"); err != nil {
		log.Exit(err)
	}

	var response *http.Response
	if response, err = snips.Create("newone"); err != nil {
		log.Exit(err)
	}

	var data []byte
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		log.Exit(err)
	}

	fmt.Printf("%v\n", string(data))

}
