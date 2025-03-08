package main

import (
	"fmt"
	"net/http"

	"gee" // 导入自定义的gee模块（gee模块是独立第三方模块，需要修改mod文件，使其从GOROOT重定向当当前目录下）
)

func main() {
	// 创建一个新的 gee.Engine 实例
	r := gee.New()

	// 注册一个处理 GET 请求的路由，路径为 "/"
	// 当访问根路径时，返回请求的 URL.Path
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	// 注册一个处理 GET 请求的路由，路径为 "/hello"
	// 当访问 "/hello" 时，返回请求头中的所有键值对
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动 HTTP 服务器，监听 9999 端口
	// 如果服务器启动失败，会返回错误
	r.Run(":9999")
}
