package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"zty/zweb"
)

/*
curl "http://localhost:9999/"
<html>
    <link rel="stylesheet" href="/assets/css/zty.css">
    <p>zty.css is loaded</p>
</html>

curl "http://localhost:9999/students"

<html>
<body>
    <p>hello, zty</p>

    <p>0: zty is 21 years old</p>

    <p>1: fmj is 23 years old</p>

</body>
</html>

curl "http://localhost:9999/date"

<html>
<body>
    <p>hello, zty</p>
    <p>Date: 2023-09-17</p>
</body>
</html>
*/

type student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := zweb.New()
	r.Use(zweb.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	stu1 := &student{Name: "zty", Age: 21}
	stu2 := &student{Name: "fmj", Age: 23}
	r.GET("/", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", zweb.H{
			"title":  "zty",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *zweb.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", zweb.H{
			"title": "zty",
			"now":   time.Date(2023, 9, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
