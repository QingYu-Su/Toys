package main

import (
	"net/http"

	"gee" // 引入 gee 框架
)

func main() {
	// 使用 gee.Default() 创建一个带有默认中间件的 Engine 实例。
	// 默认中间件包括 Logger 和 Recovery，分别用于日志记录和错误恢复。
	r := gee.Default()

	// 注册一个 GET 路由，处理根路径（"/"）的请求。
	// 当用户访问 http://localhost:9999/ 时，会触发这个处理器。
	r.GET("/", func(c *gee.Context) {
		// 使用 c.String 方法直接返回 HTTP 状态码和字符串响应。
		// 这里返回状态码 200（OK）和字符串 "Hello Geektutu\n"。
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	// 注册另一个 GET 路由，用于测试 Recovery 中间件的错误处理功能。
	// 当用户访问 http://localhost:9999/panic 时，会触发这个处理器。
	r.GET("/panic", func(c *gee.Context) {
		// 定义一个包含单个字符串的切片。
		names := []string{"geektutu"}
		// 尝试访问切片的第 100 个元素，这将导致 "index out of range" 错误。
		// 这个错误会被 Recovery 中间件捕获，并返回 500 内部服务器错误。
		c.String(http.StatusOK, names[100])
	})

	// 启动 HTTP 服务器，监听端口 9999。
	// 当服务器启动后，可以通过浏览器或工具访问 http://localhost:9999/ 和 http://localhost:9999/panic 来测试功能。
	r.Run(":9999")
}
