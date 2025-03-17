// geerpc 包提供了一个简单的 RPC 服务器实现，支持基于魔数的协议验证、
// 动态编解码器选择以及并发请求处理。
package geerpc

// 导入的包提供了以下功能：
// - encoding/json：用于处理 JSON 编解码，主要用于解析客户端发送的配置选项。
// - errors：提供错误处理功能。
// - fmt：提供格式化输出功能。
// - geerpc/codec：提供编解码器接口和实现，用于处理 RPC 请求和响应的序列化与反序列化。
// - io：提供基本的 I/O 操作接口。
// - log：提供日志记录功能，用于调试和运行时监控。
// - net：提供网络功能，用于建立客户端和服务端之间的连接。
// - reflect：提供反射功能，用于动态操作类型和值。
// - strings：提供字符串操作功能。
// - sync：提供同步原语，用于处理并发请求和响应。
// - time：提供时间相关功能，用于处理超时和延时操作。
import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
	"time"
)

// MagicNumber 是一个魔数，用于验证客户端和服务端之间的协议一致性。
const MagicNumber = 0x3bef5c

// Option 定义了 RPC 客户端和服务端的配置选项。
type Option struct {
	MagicNumber    int           // 魔数，用于验证协议一致性
	CodecType      codec.Type    // 编解码器类型，指定使用的编解码方式
	ConnectTimeout time.Duration // 连接超时时间，控制客户端与服务端建立连接的等待时间
	HandleTimeout  time.Duration // 处理超时时间，控制 RPC 请求处理的最大等待时间
}

// DefaultOption 是默认的配置选项实例，提供了一组合理的默认值。
var DefaultOption = &Option{
	MagicNumber:    MagicNumber,      // 使用全局定义的 MagicNumber 作为默认魔数
	CodecType:      codec.GobType,    // 默认使用 Gob 编解码器
	ConnectTimeout: time.Second * 10, // 默认连接超时时间为 10 秒
}

// Server 表示一个 RPC 服务器，负责管理注册的服务并处理客户端请求。
type Server struct {
	serviceMap sync.Map // 使用 sync.Map 存储服务名称到服务实例的映射，支持并发读写
}

// NewServer 创建一个新的 RPC 服务器实例。
func NewServer() *Server {
	return &Server{}
}

// DefaultServer 是默认的 RPC 服务器实例，方便直接使用。
var DefaultServer = NewServer()

// ServeConn 处理一个客户端连接，读取请求并调用对应的服务方法。
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }() // 确保连接在函数退出时关闭
	var opt Option
	// 从连接中读取客户端发送的配置选项
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}
	// 验证魔数是否一致
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	// 根据编解码类型获取对应的编解码器创建函数
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	// 使用创建的编解码器处理连接
	server.serveCodec(f(conn), &opt)
}

// invalidRequest 是一个无效请求的标记，用于发送错误响应。
var invalidRequest = struct{}{}

// serveCodec 使用指定的编解码器处理客户端请求。
func (server *Server) serveCodec(cc codec.Codec, opt *Option) {
	sending := new(sync.Mutex) // 用于同步响应发送
	wg := new(sync.WaitGroup)  // 用于等待所有请求处理完成
	for {
		// 读取客户端的请求
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil { // 如果读取请求失败且没有请求对象，则退出循环
				break
			}
			// 如果读取请求失败但有请求对象，则设置错误信息并发送错误响应
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1) // 增加等待组计数
		// 启动协程处理请求
		go server.handleRequest(cc, req, sending, wg, opt.HandleTimeout)
	}
	wg.Wait()      // 等待所有请求处理完成
	_ = cc.Close() // 关闭编解码器
}

// request 表示一个 RPC 请求，封装了处理请求所需的所有信息。
// h：请求头部信息，包含服务方法名、序列号等元数据。
// argv：请求参数的反射值，用于动态处理请求参数。
// replyv：响应值的反射值，用于存储方法调用的结果。
// mtype：方法类型，包含方法的元数据和调用逻辑。
// svc：服务实例，包含服务的元数据和方法映射。
type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
	mtype        *methodType
	svc          *service
}

// readRequestHeader 从编解码器中读取请求头部信息。
// 如果读取失败，记录错误并返回。
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil { // 从编解码器中读取请求头部
		if err != io.EOF && err != io.ErrUnexpectedEOF { // 如果不是 EOF 或意外 EOF 错误，则记录日志
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil // 返回读取到的请求头部
}

// findService 根据 serviceMethod 查找对应的 RPC 服务和方法。
// serviceMethod 的格式为 "ServiceName.MethodName"。
// 如果找到对应的 RPC 服务和方法，则返回服务实例和方法类型；否则返回错误。
func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {
	dot := strings.LastIndex(serviceMethod, ".") // 查找最后一个 '.' 的位置
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod) // 格式错误
		return
	}
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:] // 分割服务名称和服务方法名称
	svci, ok := server.serviceMap.Load(serviceName)                       // 从 serviceMap 中查找服务名称对应的 RPC 服务
	if !ok {
		err = errors.New("rpc server: can't find service " + serviceName) // 服务未找到
		return
	}
	svc = svci.(*service)          // 断言 svci 的类型为 *service
	mtype = svc.method[methodName] // 从服务的 method 映射中查找方法名称对应的 RPC 方法
	if mtype == nil {
		err = errors.New("rpc server: can't find method " + methodName) // 方法未找到
	}
	return
}

// readRequest 从编解码器中读取一个完整的 RPC 请求，并构造一个 request 对象。
// 它会依次读取请求头部和请求参数，并根据服务方法名查找对应的服务和方法。
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc) // 读取请求头部信息
	if err != nil {
		return nil, err // 如果读取头部失败，直接返回错误
	}
	req := &request{h: h}                                         // 根据头部信息创建一个 request 对象
	req.svc, req.mtype, err = server.findService(h.ServiceMethod) // 查找对应的服务和方法
	if err != nil {
		return req, err // 如果找不到服务或方法，返回错误
	}
	req.argv = req.mtype.newArgv()     // 创建请求参数的反射值
	req.replyv = req.mtype.newReplyv() // 创建响应值的反射值

	argvi := req.argv.Interface() // 获取请求参数的接口值
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface() // 如果不是指针类型，则取其地址
	}
	if err = cc.ReadBody(argvi); err != nil { // 从编解码器中读取请求参数
		log.Println("rpc server: read body err:", err) // 如果读取请求参数失败，记录错误
		return req, err
	}
	return req, nil // 返回构造好的 request 对象
}

// sendResponse 向客户端发送 RPC 响应。
// 它使用指定的编解码器写入响应头部和响应正文。
// sending 互斥锁用于确保并发写入的安全性。
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock() // 锁定发送操作，确保并发安全
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil { // 写入响应头部和正文
		log.Println("rpc server: write response error:", err) // 如果发送失败，记录错误
	}
}

// handleRequest 处理一个 RPC 请求，并在指定的超时时间内返回响应。
// 如果超时时间设置为 0，则不启用超时机制。
// called 通道用于同步方法调用完成，sent 通道用于同步响应发送完成。
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done() // 在函数返回时调用 wg.Done()，通知等待组当前请求已处理完成

	called := make(chan struct{}) // 用于同步方法调用完成
	sent := make(chan struct{})   // 用于同步响应发送完成

	// 启动一个协程执行方法调用和响应发送
	go func() {
		err := req.svc.call(req.mtype, req.argv, req.replyv) // 调用对应的服务方法
		called <- struct{}{}                                 // 通知方法调用已完成
		if err != nil {
			req.h.Error = err.Error()                               // 如果方法调用失败，设置请求头部的错误信息
			server.sendResponse(cc, req.h, invalidRequest, sending) // 发送错误响应
			sent <- struct{}{}                                      // 通知响应已发送
			return
		}
		server.sendResponse(cc, req.h, req.replyv.Interface(), sending) // 发送正常的响应
		sent <- struct{}{}                                              // 通知响应已发送
	}()

	// 如果超时时间设置为 0，等待方法调用和响应发送完成
	if timeout == 0 {
		<-called
		<-sent
		return
	}

	// 使用 select 语句处理超时逻辑
	select {
	case <-time.After(timeout): // 如果超时
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout) // 设置超时错误
		server.sendResponse(cc, req.h, invalidRequest, sending)                                    // 发送超时错误响应
	case <-called: // 如果方法调用已完成
		<-sent // 等待响应发送完成
	}
}

// Accept 在指定的监听器上接受客户端连接，并为每个连接启动一个服务协程。
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept() // 接受客户端连接
		if err != nil {
			log.Println("rpc server: accept error:", err) // 如果接受失败，记录错误
			return
		}
		go server.ServeConn(conn) // 为每个连接启动一个服务协程
	}
}

// Accept 是 DefaultServer 的 Accept 方法的简化调用。
// 它允许用户直接使用默认服务器实例接受客户端连接。
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }

// Register 将一个接收器注册为 RPC 服务。
// 接收器可以是任意类型，但必须包含符合 RPC 方法签名的导出方法。
// 如果服务已注册，则返回错误。
func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr)                                        // 创建一个服务实例
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup { // 将服务存储到 serviceMap 中
		return errors.New("rpc: service already defined: " + s.name) // 如果服务已注册，返回错误
	}
	return nil // 注册成功
}

// Register 是 DefaultServer 的 Register 方法的简化调用。
// 它允许用户直接注册服务到默认服务器实例。
func Register(rcvr interface{}) error { return DefaultServer.Register(rcvr) }
