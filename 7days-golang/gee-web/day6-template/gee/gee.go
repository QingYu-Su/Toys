package gee

import (
	"log"
	"net/http"
	"path"
	"strings"
	"text/template"
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
	*RouterGroup                     // 继承 RouterGroup 的属性和方法
	router        *router            // 路由管理器，用于管理路由
	groups        []*RouterGroup     // 存储所有路由分组
	htmlTemplates *template.Template // 用于 HTML 渲染的模板
	funcMap       template.FuncMap   // 用于 HTML 渲染的自定义函数映射
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

// Use 添加中间件到当前分组
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// createStaticHandler 创建一个静态文件处理函数
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 组合绝对路径
	absolutePath := path.Join(group.prefix, relativePath)
	// 创建一个文件服务器，并去掉绝对路径前缀
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		// 获取请求的文件路径
		file := c.Param("filepath")
		// 检查文件是否存在，以及是否有权限访问
		if _, err := fs.Open(file); err != nil {
			// 如果文件不存在或无法访问，返回 404 错误
			c.Status(http.StatusNotFound)
			return
		}
		// 使用文件服务器处理请求
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files 提供静态文件服务
func (group *RouterGroup) Static(relativePath string, root string) {
	// 创建静态文件处理函数
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	// 构造 URL 模式，支持通配符路径
	urlPattern := path.Join(relativePath, "/*filepath")
	// 注册 GET 请求的处理函数
	group.GET(urlPattern, handler)
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // 监听指定地址并启动服务器
}

// ServeHTTP 实现了 http.Handler 接口，用于处理 HTTP 请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 定义一个中间件列表，用于存储所有匹配的中间件
	var middlewares []HandlerFunc

	// 遍历所有路由分组
	for _, group := range engine.groups {
		// 检查请求路径是否以当前分组的前缀开头
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			// 如果匹配，将当前分组的中间件添加到中间件列表中
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	// 创建一个新的 Context 实例，封装请求和响应信息
	c := newContext(w, req)

	// 将中间件列表赋值给 Context 的 handlers 字段
	c.handlers = middlewares

	c.engine = engine

	// 调用路由管理器的 handle 方法，处理请求
	engine.router.handle(c)
}

// SetFuncMap 设置模板的自定义函数映射
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// LoadHTMLGlob 加载 HTML 模板文件
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
