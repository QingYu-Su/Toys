package main

import (
	"net/http"

	"gee" // 导入自定义的 gee 模块（需修改 go.mod 以确保正确导入本地模块）
)

func main() {
	r := gee.New() // 创建 Gee 框架的实例

	// 注册 GET 请求路由，访问 "/" 时返回 HTML 页面
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// 注册 GET 请求路由，访问 "/hello" 时返回格式化的字符串
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// 注册 POST 请求路由，访问 "/login" 时返回 JSON 数据
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 启动服务器，监听 9999 端口
	r.Run(":9999")
}
