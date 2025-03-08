package gee

import (
	"net/http"
	"strings"
)

// router 是一个路由管理器，用于存储和匹配HTTP请求的路由
type router struct {
	roots    map[string]*node       // 每个HTTP方法（如GET、POST）对应一个前缀树的根节点
	handlers map[string]HandlerFunc // 存储HTTP方法和路径模式对应的处理函数
}

// newRouter 创建一个新的路由管理器
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),       // 初始化根节点映射
		handlers: make(map[string]HandlerFunc), // 初始化处理函数映射
	}
}

// parsePattern 将路径模式拆分成多个片段
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/") // 按 '/' 分割路径模式

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" { // 跳过空字符串
			parts = append(parts, item)
			if item[0] == '*' { // 如果遇到通配符 '*'，停止拆分
				break
			}
		}
	}
	return parts
}

// addRoute 向路由管理器中添加一个新的路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) // 拆分路径模式

	key := method + "-" + pattern // 生成唯一键，用于存储处理函数
	_, ok := r.roots[method]      // 检查是否存在该HTTP方法的根节点
	if !ok {
		r.roots[method] = &node{} // 如果不存在，创建一个新的根节点
	}

	r.roots[method].insert(pattern, parts, 0) // 将路径模式插入到前缀树中
	r.handlers[key] = handler                 // 将处理函数存储到映射中
}

// getRoute 根据HTTP方法和请求路径查找匹配的路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) // 拆分请求路径
	params := make(map[string]string) // 用于存储路径参数
	root, ok := r.roots[method]       // 获取对应HTTP方法的根节点

	if !ok { // 如果没有找到根节点，返回 nil
		return nil, nil
	}

	n := root.search(searchParts, 0) // 在前缀树中查找匹配的节点

	if n != nil { // 如果找到匹配的节点
		parts := parsePattern(n.pattern) // 拆分匹配的路径模式
		for index, part := range parts {
			if part[0] == ':' { // 如果是动态参数（以 ':' 开头）
				params[part[1:]] = searchParts[index] // 提取参数值
			}
			if part[0] == '*' && len(part) > 1 { // 如果是通配符参数（以 '*' 开头）
				params[part[1:]] = strings.Join(searchParts[index:], "/") // 提取通配符参数的值
				break
			}
		}
		return n, params // 返回匹配的节点和参数
	}

	return nil, nil // 如果没有找到匹配的节点，返回 nil
}

// getRoutes 获取指定HTTP方法的所有路由
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method] // 获取对应HTTP方法的根节点

	if !ok { // 如果没有找到根节点，返回 nil
		return nil
	}

	nodes := make([]*node, 0) // 创建一个空的节点列表
	root.travel(&nodes)       // 遍历前缀树，将所有节点添加到列表中
	return nodes
}

// handle 处理HTTP请求
func (r *router) handle(c *Context) {
	// 调用 getRoute 方法，根据请求的方法和路径查找匹配的路由
	n, params := r.getRoute(c.Method, c.Path)

	// 如果找到匹配的路由
	if n != nil {
		// 将路径参数存储到 Context 的 Params 中
		c.Params = params
		// 生成唯一键，格式为 "HTTP方法-路径模式"
		key := c.Method + "-" + n.pattern
		// 将对应的处理函数添加到 Context 的处理函数链中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		// 如果没有找到匹配的路由，添加一个默认的 404 错误处理函数
		c.handlers = append(c.handlers, func(c *Context) {
			// 返回 404 错误
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	// 调用 Context 的 Next 方法，继续处理请求
	c.Next()
}
