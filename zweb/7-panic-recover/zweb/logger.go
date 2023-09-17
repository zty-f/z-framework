package zweb

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		//记录日志
		log.Printf("[%d] %s expend %v", c.StatusCode, c.Method+"-"+c.Path, time.Since(t))
	}
}
