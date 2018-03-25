package main

import (
	"fmt"
	"net/http"
)

var endpoint = "http://localhost:18888"

// List 3-6 (P73)
func main() {
	resp, err := http.Head(endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
}
