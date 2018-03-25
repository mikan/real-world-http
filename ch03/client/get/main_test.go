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
	host = ts.URL
	main()
}
