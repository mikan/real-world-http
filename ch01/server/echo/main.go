package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// List 1-1 (P4)
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprint(w, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	fmt.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	fmt.Println(httpServer.ListenAndServe())
}
