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
	log.Println("Starting Server")
	address := "127.0.0.1:3000"
	snips := NewSnipsCollection()

	snips.Add("first post!")
	snips.Add("me too")

	rest.Resource("snips", snips)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalln(err)
	}
}

func client() {
	log.Println("Starting Client")
	var snips *rest.Client
	var err os.Error

	if snips, err = rest.NewClient("http://127.0.0.1:3000/snips/"); err != nil {
		log.Fatalln(err)
	}

	// Create a new snip
	var response *http.Response
	if response, err = snips.Create("newone"); err != nil {
		log.Fatalln(err)
	}

	var id string
	if id, err = snips.IdFromURL(response.Header.Get("Location")); err != nil {
		log.Fatalln(err)
	}

	// Update the snip
	if response, err = snips.Update(id, "updated"); err != nil {
		log.Fatalln(err)
	}


	// Get the updated snip
	if response, err = snips.Find(id); err != nil {
		log.Fatalln(err)
	}

	var data []byte
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", string(data))

	// Delete the created snip
	if response, err = snips.Delete(id); err != nil {
		log.Fatalln(err)
	}

}
