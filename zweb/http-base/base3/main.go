package main

import (
	"example/zweb"
	"fmt"
	"net/http"
)

func main() {
	engine := zweb.New()

	//curl http://localhost:8080/
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	// curl -X POST http://localhost:8080/hello
	engine.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		fmt.Println("run err:", err)
		return
	}
}
