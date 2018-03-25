package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

var host = "http://localhost:18888"

func main() {
	resp, err := http.Get(host)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
