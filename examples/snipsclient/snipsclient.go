package main

import (
	"fmt"
	"log"
	"http"
	"io/ioutil"
	"os"
	"github.com/nathankerr/rest.go"
)

func main() {
	var snips *rest.Client
	var err os.Error

	if snips, err = rest.NewClient("http://127.0.0.1:3000/snips/"); err != nil {
		log.Exit(err)
	}

	var response *http.Response
	if response, err = snips.Index(); err != nil {
		log.Exit(err)
	}

	var data []byte
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		log.Exit(err)
	}

	fmt.Printf("%v\n", string(data))
}
