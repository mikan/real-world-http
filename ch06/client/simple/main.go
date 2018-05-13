package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}

// panic: Get https://localhost:18443: x509: certificate signed by unknown authority
//
// goroutine 1 [running]:
// main.main()
// <GOPATH>/src/github.com/mikan/real-world-http/ch06/client/simple/main.go:12 +0x195
// exit status 2
