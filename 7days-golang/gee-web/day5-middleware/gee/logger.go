package gee

import (
	"log"
	"time"
)

// Logger 返回一个 HandlerFunc，用于记录请求的处理时间和状态码
func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now() // 记录请求开始的时间
		c.Next()        // 调用下一个处理函数，继续处理请求
		// 记录请求的处理时间、状态码和请求 URI
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
