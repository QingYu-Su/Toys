package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义了一个处理HTTP请求的函数类型
// 它接收 http.ResponseWriter 和 *http.Request 作为参数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 是框架的核心结构，包含一个路由表（router）
// router 是一个映射，键是方法（GET/POST等）和路径的组合，值是对应的处理函数
type Engine struct {
	router map[string]HandlerFunc
}

// New 是 Engine 的构造函数，用于创建一个新的 Engine 实例
// 初始化时创建一个空的 router 映射
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 是一个内部方法，用于将路由规则添加到路由表中
// method: HTTP 方法（如 GET、POST）
// pattern: URL 路径（如 "/hello"）
// handler: 处理该路由的函数
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // 构造路由键，格式为 "方法-路径"
	engine.router[key] = handler  // 将处理函数与路由键关联
}

// GET 是一个便捷方法，用于注册处理 GET 请求的路由
// pattern: URL 路径
// handler: 处理该路由的函数
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler) // 调用 addRoute 方法，添加 GET 路由
}

// POST 是一个便捷方法，用于注册处理 POST 请求的路由
// pattern: URL 路径
// handler: 处理该路由的函数
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler) // 调用 addRoute 方法，添加 POST 路由
}

// Run 启动 HTTP 服务器，监听指定的地址
// addr: 服务器监听的地址（如 ":8080"）
// 返回一个 error，如果服务器启动失败，则返回错误信息
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // 启动服务器，并将 Engine 实例作为处理器
}

// ServeHTTP 实现了 http.Handler 接口，用于处理 HTTP 请求
// 根据请求的方法和路径查找对应的处理函数，如果找到则调用，否则返回 404 错误
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path // 构造当前请求的路由键
	if handler, ok := engine.router[key]; ok {
		handler(w, req) // 如果找到对应的处理函数，则调用它
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL) // 如果未找到路由，返回 404 错误
	}
}
