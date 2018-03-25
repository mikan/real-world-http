package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var endpoint = "http://localhost:18888"

// List 3-4 (P72)
func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get(endpoint + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
