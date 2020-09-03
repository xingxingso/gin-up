package ginseng

import (
	"fmt"
	"log"
	"net/http"
)

func Run(addr string) {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}