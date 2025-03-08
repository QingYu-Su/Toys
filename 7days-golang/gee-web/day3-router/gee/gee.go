package gee

import "net/http"

// HandlerFunc 定义请求处理函数类型
type HandlerFunc func(*Context)

// Engine 是框架的核心结构体，实现了 http.Handler 接口
type Engine struct {
	router *router // 路由管理器
}

// New 创建并返回一个新的 Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加路由规则到 router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 注册一个 GET 请求的路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 注册一个 POST 请求的路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP 服务器，监听指定地址
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现 http.Handler 接口，用于处理 HTTP 请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req) // 创建请求上下文
	engine.router.handle(c) // 通过 router 处理请求
}
