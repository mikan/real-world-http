package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var endpoint = "http://localhost:18888"

// List 3-1 (P58)
func main() {
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
}
