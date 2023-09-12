package main

import (
	"net/http"
	"zty/zweb"
)

/*
curl -i http://localhost:9999/
HTTP/1.1 200 OK
Content-Type: text/html
Date: Tue, 12 Sep 2023 07:06:34 GMT
Content-Length: 18

<h1>Hello zty</h1>%

curl "http://localhost:9999/hello?name=zty"
hello zty, you're at /hello

curl "http://localhost:9999/hello/zty"
hello zty, you're at /hello/zty

curl "http://localhost:9999/assets/css/zty.css"
{"filepath":"css/zty.css"}

curl "http://localhost:9999/xxx"
404 Not Found: /xxx
*/

func main() {
	r := zweb.New()
	r.GET("/", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello zty</h1>")
	})

	r.GET("/hello", func(c *zweb.Context) {
		// expect /hello?name=zty
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *zweb.Context) {
		// expect /hello/zty
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *zweb.Context) {
		c.JSON(http.StatusOK, zweb.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
