package main

import (
	"log"
	"net/http"
	"time"

	"gee" // 导入自定义的 gee 模块（需修改 go.mod 以确保正确导入本地模块）
)

// onlyForV2 定义了一个只适用于 v2 分组的中间件
func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()                      // 记录请求开始的时间
		c.Fail(500, "Internal Server Error") // 调用 Fail 方法，返回 500 错误
		// 记录请求的处理时间、状态码和请求 URI
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New() // 创建一个新的 gee.Engine 实例

	// 使用全局中间件 Logger
	r.Use(gee.Logger())

	// 定义一个全局路由
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>") // 返回 HTML 内容
	})

	// 创建一个路由分组 "/v2"
	v2 := r.Group("/v2")
	// 为 v2 分组添加中间件 onlyForV2
	v2.Use(onlyForV2())

	// 定义 v2 分组下的路由
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// 从路径参数中获取 "name" 的值
			// 例如，访问 "/v2/hello/geektutu"
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	// 启动服务器，监听 9999 端口
	r.Run(":9999")
}
