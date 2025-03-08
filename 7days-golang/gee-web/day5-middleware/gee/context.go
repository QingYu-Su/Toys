package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是一个快捷方式，定义 map[string]interface{} 类型，方便 JSON 数据返回
type H map[string]interface{}

// Context 封装了 HTTP 请求的上下文信息
type Context struct {
	Writer     http.ResponseWriter // HTTP 响应写入器
	Req        *http.Request       // HTTP 请求
	Path       string              // 请求路径
	Method     string              // 请求方法 (GET, POST 等)
	Params     map[string]string   // 路径参数
	StatusCode int                 // HTTP 响应状态码
	handlers   []HandlerFunc       // 处理函数链
	index      int                 // 当前处理函数的索引
}

// newContext 创建一个新的 Context 实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1, // 初始化索引为 -1，表示尚未开始处理
	}
}

// Fail 用于处理失败情况，返回一个 JSON 响应
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)       // 设置索引为处理函数链的长度，跳过后续处理
	c.JSON(code, H{"message": err}) // 返回一个 JSON 响应，包含错误信息
}

// PostForm 获取 POST 表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key) // 使用 http.Request 的 FormValue 方法获取表单参数
}

// Query 获取 URL 查询参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key) // 使用 http.Request 的 URL.Query().Get 方法获取查询参数
}

// Status 设置 HTTP 响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code) // 调用 http.ResponseWriter 的 WriteHeader 方法设置状态码
}

// SetHeader 设置 HTTP 响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value) // 调用 http.ResponseWriter 的 Header().Set 方法设置响应头
}

// String 以文本格式返回响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")              // 设置响应内容类型为 text/plain
	c.Status(code)                                         // 设置 HTTP 状态码
	c.Writer.Write([]byte(fmt.Sprintf(format, values...))) // 写入格式化字符串
}

// JSON 以 JSON 格式返回响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json") // 设置响应内容类型为 application/json
	c.Status(code)                                  // 设置 HTTP 状态码
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil { // JSON 序列化并写入响应
		http.Error(c.Writer, err.Error(), 500) // 处理编码错误
	}
}

// Data 以字节流方式返回响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)       // 设置 HTTP 状态码
	c.Writer.Write(data) // 写入字节数据
}

// HTML 以 HTML 格式返回响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html") // 设置响应内容类型为 text/html
	c.Status(code)                           // 设置 HTTP 状态码
	c.Writer.Write([]byte(html))             // 写入 HTML 内容
}

// Param 获取路径参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key] // 从 Params 映射中获取路径参数
	return value
}

// Next 调用下一个处理函数
func (c *Context) Next() {
	c.index++                      // 增加索引
	s := len(c.handlers)           // 获取处理函数链的长度
	for ; c.index < s; c.index++ { // 遍历处理函数链
		c.handlers[c.index](c) // 调用当前索引的处理函数
	}
}
