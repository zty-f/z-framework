package main

import (
	"net/http"
	"zty/zweb"
)

/*
zweb 是一个轻量级的web框架，它提供了一些常用的功能，例如：路由、中间件、模板、静态文件、Session、JSON

curl -i http://localhost:8080/
HTTP/1.1 200 OK
Content-Type: text/html
Date: Sun, 10 Sep 2023 08:31:25 GMT
Content-Length: 21

<h1>Hello, zweb!</h1>%

curl "http://localhost:8080/hello?name=zty"
Hello zty, you are in zweb: /hello!%

curl "http://localhost:8080/login" -X POST -d 'username=zty&password=1234'
{"password":"1234","username":"zty"}
*/

func main() {
	r := zweb.New()

	r.GET("/", func(c *zweb.Context) {
		c.HTML(200, "<h1>Hello, zweb!</h1>")
	})

	r.GET("/hello", func(c *zweb.Context) {
		c.String(200, "Hello %s, you are in zweb: %s!", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *zweb.Context) {
		c.JSON(http.StatusOK, zweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
