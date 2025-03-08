package main

import (
	"net/http"

	"gee" // 导入自定义的 gee 模块（需修改 go.mod 以确保正确导入本地模块）
)

func main() {
	// 创建一个新的 gee.Engine 实例
	r := gee.New()

	// 定义一个 GET 路由，处理根路径 "/"
	r.GET("/index", func(c *gee.Context) {
		// 返回一个简单的 HTML 内容
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	// 创建一个路由分组 "/v1"
	v1 := r.Group("/v1")
	{
		// 定义一个 GET 路由，处理 "/v1/"
		v1.GET("/", func(c *gee.Context) {
			// 返回一个简单的 HTML 内容
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		// 定义一个 GET 路由，处理 "/v1/hello"
		v1.GET("/hello", func(c *gee.Context) {
			// 从 URL 查询参数中获取 "name" 的值
			// 例如，访问 "/v1/hello?name=geektutu"
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	// 创建另一个路由分组 "/v2"
	v2 := r.Group("/v2")
	{
		// 定义一个 GET 路由，处理 "/v2/hello/:name"
		v2.GET("/hello/:name", func(c *gee.Context) {
			// 从路径参数中获取 "name" 的值
			// 例如，访问 "/v2/hello/geektutu"
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		// 定义一个 POST 路由，处理 "/v2/login"
		v2.POST("/login", func(c *gee.Context) {
			// 从 POST 表单数据中获取 "username" 和 "password"
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	// 启动 HTTP 服务器，监听 9999 端口
	r.Run(":9999")
}
