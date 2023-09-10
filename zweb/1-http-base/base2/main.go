package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the handler for all requests
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Welcome to the home page! %s\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL)
	}
}

func main() {
	engin := new(Engine)

	log.Fatal(http.ListenAndServe(":8080", engin))
}
