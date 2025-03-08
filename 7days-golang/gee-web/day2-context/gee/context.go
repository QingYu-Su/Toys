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
	StatusCode int                 // HTTP 响应状态码
}

// newContext 创建一个新的 Context 实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取 POST 表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取 URL 查询参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置 HTTP 响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置 HTTP 响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
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
