package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestToMain(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("must use %s method", http.MethodGet)
		}
		w.Write([]byte("Hello, World!"))
	}))
	endpoint = ts.URL
	main()
}

func Example() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	endpoint = ts.URL
	main()
	// Output:
	// Hello, World!
	// Status: 200 OK
	// StatusCode: 200
}
