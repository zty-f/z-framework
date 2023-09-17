package main

import (
	"log"
	"net/http"
	"time"
	"zty/zweb"
)

/*
curl "http://localhost:9999/"
<h1>Hello zty</h1>%

log: [200] GET-/ expend 1.75µs

curl "http://localhost:9999/v1/hello/zty"
{"message":"Internal Server Error"}

log：[500] /v1/hello/zty in 181.083µs for group v1
     [500] GET-/v1/hello/zty expend 203.166µs
*/

func onlyForV1() zweb.HandlerFunc {
	return func(c *zweb.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v1", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := zweb.New()
	r.Use(zweb.Logger())
	r.GET("/", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello zty</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(onlyForV1())
	v1.GET("/hello/:name", func(c *zweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Run(":9999")
}
