package codec

import "io"

// Header 定义了 RPC 请求或响应的头部信息
type Header struct {
	ServiceMethod string // 服务方法名，表示要调用的服务和方法
	Seq           uint64 // 序列号，用于匹配请求和响应
	Error         string // 错误信息，用于表示处理过程中出现的错误
}

// Codec 定义了编解码器的接口，用于处理 RPC 的编解码操作
type Codec interface {
	io.Closer // 实现 io.Closer 接口，用于关闭底层的 io 流
	// 读取请求或响应的头部信息
	ReadHeader(*Header) error
	// 读取请求或响应的正文内容
	ReadBody(interface{}) error
	// 写入请求或响应的头部和正文内容
	Write(*Header, interface{}) error
}

// NewCodecFunc 定义了创建 Codec 实例的函数类型
type NewCodecFunc func(io.ReadWriteCloser) Codec

// Type 定义了编解码器的类型，用于区分不同的编解码方式
type Type string

// 定义了两种常见的编解码类型
const (
	GobType  Type = "application/gob"  // Gob 编解码类型
	JsonType Type = "application/json" // JSON 编解码类型
)

// NewCodecFuncMap 用于存储不同编解码类型对应的创建函数
var NewCodecFuncMap map[Type]NewCodecFunc

// 初始化 NewCodecFuncMap，注册 Gob 编解码器的创建函数
func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec // 注册 Gob 编解码器的创建函数
}
