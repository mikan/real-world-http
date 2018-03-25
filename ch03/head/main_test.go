package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestToMain(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Fatalf("must use %s method", http.MethodHead)
		}
	}))
	endpoint = ts.URL
	main()
}

func Example() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Sun, 25 Mar 2018 12:00:00 GMT")
	}))
	endpoint = ts.URL
	main()
	// Output:
	// Status: 200 OK
	// Headers: map[Date:[Sun, 25 Mar 2018 12:00:00 GMT]]
}
