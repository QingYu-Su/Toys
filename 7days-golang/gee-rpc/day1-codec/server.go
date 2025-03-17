package geerpc

import (
	"encoding/json" // 用于处理 JSON 编解码
	"fmt"           // 用于格式化输出
	"geerpc/codec"  // 引入自定义的 codec 包，用于处理 RPC 编解码
	"io"            // 提供基本的 I/O 操作接口
	"log"           // 用于记录日志信息
	"net"           // 提供网络相关功能
	"reflect"       // 提供反射功能，用于动态处理类型和值
	"sync"          // 提供同步原语，用于并发控制
)

// MagicNumber 是一个魔数，用于验证客户端和服务器之间的协议一致性
const MagicNumber = 0x3bef5c

// Option 定义了 RPC 服务器的配置选项
type Option struct {
	MagicNumber int        // 魔数，用于验证协议一致性
	CodecType   codec.Type // 编解码类型，指定使用的编解码器
}

// DefaultOption 是默认的 RPC 服务器配置选项
var DefaultOption = &Option{
	MagicNumber: MagicNumber,   // 默认魔数
	CodecType:   codec.GobType, // 默认使用 Gob 编解码器
}

// Server 是一个 RPC 服务器
type Server struct{}

// NewServer 创建一个新的 RPC 服务器实例
func NewServer() *Server {
	return &Server{}
}

// DefaultServer 是默认的 RPC 服务器实例
var DefaultServer = NewServer()

// Accept 在指定的监听器上接受客户端连接，并为每个连接启动一个服务协程
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept() // 接受客户端连接
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn) // 为每个连接启动一个服务协程
	}
}

// Accept 是 DefaultServer 的 Accept 方法的简化调用
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

// ServeConn 为一个客户端连接提供 RPC 服务
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }() // 确保连接在函数退出时关闭
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil { // 从连接中读取客户端发送的配置选项
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber { // 验证魔数是否一致
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType] // 根据编解码类型获取对应的编解码器创建函数
	if f == nil {                             // 如果编解码类型无效
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn)) // 使用创建的编解码器为连接提供服务
}

// invalidRequest 是一个无效请求的标记
var invalidRequest = struct{}{}

// serveCodec 使用指定的编解码器为连接提供 RPC 服务
func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex) // 用于同步响应发送
	wg := new(sync.WaitGroup)  // 用于等待所有请求处理完成
	for {
		req, err := server.readRequest(cc) // 读取客户端的请求
		if err != nil {
			if req == nil { // 如果读取请求失败且没有请求对象，则退出循环
				break
			}
			req.h.Error = err.Error()                               // 如果读取请求失败但有请求对象，则设置错误信息
			server.sendResponse(cc, req.h, invalidRequest, sending) // 发送错误响应
			continue
		}
		wg.Add(1)                                     // 增加等待组计数
		go server.handleRequest(cc, req, sending, wg) // 启动协程处理请求
	}
	wg.Wait()      // 等待所有请求处理完成
	_ = cc.Close() // 关闭编解码器
}

// request 表示一个 RPC 请求
type request struct {
	h            *codec.Header // 请求头部信息
	argv, replyv reflect.Value // 请求参数和响应值的反射值
}

// readRequestHeader 从编解码器中读取请求头部信息
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil { // 从编解码器中读取头部信息
		if err != io.EOF && err != io.ErrUnexpectedEOF { // 如果不是 EOF 或意外 EOF 错误，则记录日志
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

// readRequest 从编解码器中读取一个完整的 RPC 请求
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc) // 读取请求头部信息
	if err != nil {
		return nil, err
	}
	req := &request{h: h}
	req.argv = reflect.New(reflect.TypeOf(""))               // 创建请求参数的反射值
	if err = cc.ReadBody(req.argv.Interface()); err != nil { // 从编解码器中读取请求参数
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

// sendResponse 向客户端发送 RPC 响应
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock() // 锁定发送操作
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil { // 使用编解码器发送响应
		log.Println("rpc server: write response error:", err)
	}
}

// handleRequest 处理一个 RPC 请求
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()                                                        // 请求处理完成后减少等待组计数
	log.Println(req.h, req.argv.Elem())                                    // 记录请求头部信息和请求参数
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq)) // 创建响应值
	server.sendResponse(cc, req.h, req.replyv.Interface(), sending)        // 发送响应
}
