package main

import (
	"net/http"

	"gee" // 导入自定义的 gee 模块（需修改 go.mod 以确保正确导入本地模块）
)

func main() {
	r := gee.New() // 创建一个新的 gee 路由管理器
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>") // 返回 HTML 内容
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path) // 返回字符串内容，使用 URL 查询参数
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path) // 返回字符串内容，使用路径参数
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")}) // 返回 JSON 内容，使用通配符参数
	})

	r.Run(":9999") // 启动服务器，监听 9999 端口
}
