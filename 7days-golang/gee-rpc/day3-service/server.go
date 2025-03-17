package geerpc

import (
	"encoding/json" // 用于处理 JSON 编解码
	"errors"        // 用于格式化输出
	"geerpc/codec"  // 引入自定义的 codec 包，用于处理 RPC 编解码
	"io"            // 提供基本的 I/O 操作接口
	"log"           // 用于记录日志信息
	"net"           // 提供网络相关功能
	"reflect"       // 提供反射功能，用于动态处理类型和值
	"strings"
	"sync" // 提供同步原语，用于并发控制
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

// Server 表示一个 RPC 服务器，负责管理注册的服务。
type Server struct {
	serviceMap sync.Map // 使用 sync.Map 存储服务名称到服务实例的映射，支持并发读写
}

// Register 将一个接收器注册为 RPC 服务。
// 接收器可以是任意类型，但必须包含符合 RPC 方法签名的导出方法。
func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr) // 创建一个服务实例
	// 使用 sync.Map 的 LoadOrStore 方法将服务名称映射到服务实例：
	// - 如果服务名称已存在，则返回已存在的服务实例和 true。
	// - 如果服务名称不存在，则存储新的服务实例并返回 nil 和 false。
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup {
		return errors.New("rpc: service already defined: " + s.name) // 如果服务已注册，返回错误
	}
	return nil // 注册成功
}

// Register 是 DefaultServer 的 Register 方法的简化调用。
// 它允许用户直接注册服务到默认服务器实例。
func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr) // 调用 DefaultServer 的 Register 方法
}

// findService 根据 serviceMethod 查找对应的 RPC 服务和方法。
// serviceMethod 的格式为 "ServiceName.MethodName"。
// 如果找到对应的 RPC 服务和方法，则返回服务实例和方法类型；
// 如果查找失败，则返回相应的错误。
func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {
	// 查找最后一个 '.' 的位置，用于分割服务名称和方法名称
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		// 如果没有找到 '.'，说明 serviceMethod 格式不正确
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}
	// 分割 serviceMethod 为服务名称和服务方法名称
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:]

	// 从 serviceMap 中查找服务名称对应的 RPC 服务
	svci, ok := server.serviceMap.Load(serviceName)
	if !ok {
		// 如果服务名称不存在，返回错误
		err = errors.New("rpc server: can't find service " + serviceName)
		return
	}
	// 断言 svci 的类型为 *service
	svc = svci.(*service)

	// 从服务的 method 映射中查找方法名称对应的 RPC 方法
	mtype = svc.method[methodName]
	if mtype == nil {
		// 如果方法名称不存在，返回错误
		err = errors.New("rpc server: can't find method " + methodName)
	}
	return
}

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

// request 表示一个 RPC 请求，封装了处理 RPC 调用所需的所有信息。
type request struct {
	h            *codec.Header // 请求头部信息，包含服务方法名、序列号等元数据
	argv, replyv reflect.Value // 请求参数和响应值的反射值，用于动态处理请求和响应
	mtype        *methodType   // 指向该请求对应的方法类型，包含方法的元数据和调用逻辑
	svc          *service      // 指向该请求对应的服务实例，包含服务的元数据和方法映射
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

// readRequest 从指定的编解码器中读取一个完整的 RPC 请求，并构造一个 request 对象。
// 它会依次读取请求头部和请求参数，并根据服务方法名查找对应的服务和方法。
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	// 首先读取请求头部信息
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err // 如果读取头部失败，直接返回错误
	}

	// 根据头部信息创建一个 request 对象
	req := &request{h: h}

	// 根据服务方法名查找对应的服务和方法
	req.svc, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err // 如果找不到服务或方法，返回错误
	}

	// 根据方法类型创建请求参数和响应值的反射值
	req.argv = req.mtype.newArgv()
	req.replyv = req.mtype.newReplyv()

	// 确保 argv 是一个指针类型，因为 ReadBody 方法需要一个指针作为参数
	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface() // 如果不是指针类型，则取其地址
	}

	// 从编解码器中读取请求参数
	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err:", err) // 如果读取请求参数失败，记录错误
		return req, err
	}

	// 如果一切正常，返回构造好的 request 对象
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

// handleRequest 处理一个 RPC 请求，并将响应发送回客户端。
// 它调用对应的服务方法，处理请求参数，并将结果写入客户端。
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数返回时调用 wg.Done()，通知等待组当前请求已处理完成

	// 调用对应的服务方法
	err := req.svc.call(req.mtype, req.argv, req.replyv)
	if err != nil {
		// 如果方法调用失败，设置请求头部的错误信息
		req.h.Error = err.Error()
		// 发送错误响应
		server.sendResponse(cc, req.h, invalidRequest, sending)
		return
	}

	// 如果方法调用成功，发送正常的响应
	server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
}
