package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义了请求处理函数的类型，每个处理函数接收一个 *Context 参数
type HandlerFunc func(*Context)

// RouterGroup 定义了路由分组的结构
type RouterGroup struct {
	prefix      string        // 当前分组的前缀路径
	middlewares []HandlerFunc // 当前分组的中间件列表
	engine      *Engine       // 所有分组共享同一个 Engine 实例
}

// Engine 是框架的核心结构，实现了 http.Handler 接口
type Engine struct {
	*RouterGroup                // 继承 RouterGroup 的属性和方法
	router       *router        // 路由管理器，用于管理路由
	groups       []*RouterGroup // 存储所有路由分组
}

// New 是 Engine 的构造函数，用于创建一个新的 Engine 实例
func New() *Engine {
	engine := &Engine{router: newRouter()}             // 创建一个新的 Engine 实例，初始化 router
	engine.RouterGroup = &RouterGroup{engine: engine}  // 初始化根路由分组
	engine.groups = []*RouterGroup{engine.RouterGroup} // 将根分组添加到分组列表
	return engine
}

// Group 创建一个新的路由分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine // 获取当前分组的 Engine 实例
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, // 新分组的前缀是当前分组前缀和传入前缀的组合
		engine: engine,                // 新分组共享同一个 Engine 实例
	}
	engine.groups = append(engine.groups, newGroup) // 将新分组添加到 Engine 的分组列表
	return newGroup
}

// addRoute 添加一个新的路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp                         // 组合分组前缀和路由路径，生成完整的路径模式
	log.Printf("Route %4s - %s", method, pattern)          // 打印路由信息，便于调试
	group.engine.router.addRoute(method, pattern, handler) // 将路由添加到路由管理器
}

// GET 添加一个 GET 请求路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler) // 调用 addRoute 方法，指定 HTTP 方法为 GET
}

// POST 添加一个 POST 请求路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler) // 调用 addRoute 方法，指定 HTTP 方法为 POST
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // 监听指定地址并启动服务器
}

// ServeHTTP 实现了 http.Handler 接口，用于处理 HTTP 请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req) // 创建一个新的 Context 实例，封装请求和响应信息
	engine.router.handle(c) // 调用路由管理器的 handle 方法，处理请求
}
