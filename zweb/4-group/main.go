package main

import (
	"net/http"
	"zty/zweb"
)

/*
curl "http://localhost:9999/index"
<h1>Hello zty</h1>%

curl "http://localhost:9999/v1/"
hello , you're at /v1/

curl "http://localhost:9999/v1/hello"
hello , you're at /v1/hello

curl "http://localhost:9999/v2/hello/zty"
hello zty, you're at /v2/hello/zty

curl -X POST "http://localhost:9999/v2/login" -d 'username=zty&password=123456'
{"password":"123456","username":"zty"}

curl "http://localhost:9999/v2/hele"
404 Not Found: /v2/hele
*/

func main() {
	r := zweb.New()
	r.GET("/index", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello zty</h1>")
	})

	v1 := r.Group("/v1")

	v1.GET("/", func(c *zweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v1.GET("/hello", func(c *zweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	v2 := r.Group("/v2")
	v2.GET("/hello/:name", func(c *zweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v2.POST("/login", func(c *zweb.Context) {
		c.JSON(http.StatusOK, zweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
