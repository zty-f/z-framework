package main

import (
	"net/http"
	"zty/zweb"
)

/*
curl "http://localhost:9999/"
hello zty!

curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}

log:
runtime error: index out of range [100] with length 1
Traceback:
        /Users/xwx/go-1.20/go1.18/src/runtime/panic.go:839
        /Users/xwx/go-1.20/go1.18/src/runtime/panic.go:89
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/main.go:40
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/context.go:43
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/recovery.go:36
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/context.go:43
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/logger.go:13
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/context.go:43
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/router.go:86
        /Users/xwx/Downloads/项目/z-framework/zweb/7-panic-recover/zweb/zweb.go:120
        /Users/xwx/go-1.20/go1.18/src/net/http/server.go:2917
        /Users/xwx/go-1.20/go1.18/src/net/http/server.go:1967
        /Users/xwx/go-1.20/go1.18/src/runtime/asm_arm64.s:1260

*/

func main() {
	r := zweb.Default()
	r.GET("/", func(c *zweb.Context) {
		c.String(http.StatusOK, "hello zty!")
	})
	r.GET("/panic", func(c *zweb.Context) {
		names := []string{"zty"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
