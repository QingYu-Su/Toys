package gee

import (
	"log"
	"net/http"
)

// router 结构体用于存储路由映射关系
type router struct {
	handlers map[string]HandlerFunc // 存储请求方法 + 路径作为键，对应的处理函数
}

// newRouter 创建一个新的 router 实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute 添加路由到路由表中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern) // 记录路由信息
	key := method + "-" + pattern                 // 生成路由的唯一标识符
	r.handlers[key] = handler                     // 存储到路由映射表
}

// handle 处理请求，匹配路由并执行对应的处理函数
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path // 生成查询的路由键
	if handler, ok := r.handlers[key]; ok {
		handler(c) // 调用对应的处理函数
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) // 处理 404 错误
	}
}
