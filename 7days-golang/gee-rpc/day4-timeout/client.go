package geerpc

import (
	"context"       // 提供上下文管理功能，用于超时和取消操作
	"encoding/json" // 提供 JSON 编解码功能，用于处理客户端和服务端之间的配置信息
	"errors"        // 提供错误处理功能
	"fmt"           // 提供格式化输出功能
	"geerpc/codec"  // 提供编解码器接口和实现，用于处理 RPC 请求和响应的序列化和反序列化
	"io"            // 提供基本的 I/O 操作接口，确保 Client 实现 io.Closer 接口
	"log"           // 提供日志记录功能，用于调试和运行时监控
	"net"           // 提供网络功能，用于建立客户端和服务端之间的连接
	"sync"          // 提供同步原语，用于处理并发请求和响应
	"time"          // 提供时间相关功能，用于处理超时和延时操作
)

// Call 表示一个 RPC 调用的上下文，包含请求和响应的所有信息。
type Call struct {
	Seq           uint64      // 请求的唯一序列号，用于匹配请求和响应
	ServiceMethod string      // 要调用的服务方法名，格式为 "Service.Method"
	Args          interface{} // 请求参数，可以是任意类型
	Reply         interface{} // 响应结果，可以是任意类型
	Error         error       // 调用过程中发生的错误
	Done          chan *Call  // 用于通知调用完成的通道，当调用完成时，会向该通道发送自身
}

// done 将当前 Call 实例发送到 Done 通道，表示调用已完成。
func (call *Call) done() {
	call.Done <- call
}

// Client 是一个 RPC 客户端，用于管理与服务器的连接，并发起 RPC 调用。
type Client struct {
	cc       codec.Codec      // 编解码器，用于处理 RPC 请求和响应的编解码
	opt      *Option          // 客户端配置选项，例如编解码类型和超时时间
	sending  sync.Mutex       // 互斥锁，用于同步发送请求的过程，确保请求顺序正确
	header   codec.Header     // 请求头部信息，包含服务方法名、序列号等
	mu       sync.Mutex       // 互斥锁，用于同步访问共享资源（如 pending 和 seq）
	seq      uint64           // 当前请求的序列号，每次调用递增
	pending  map[uint64]*Call // 存储尚未完成的调用，键为序列号，值为 Call 实例
	closing  bool             // 标记客户端是否正在关闭
	shutdown bool             // 标记客户端是否已完全关闭
}

// 确保 Client 实现了 io.Closer 接口，以便可以关闭底层连接。
var _ io.Closer = (*Client)(nil)

// ErrShutdown 是一个错误，表示客户端连接已被关闭。
var ErrShutdown = errors.New("connection is shut down")

// Close 关闭客户端连接。如果客户端已经关闭，则返回 ErrShutdown。
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing { // 如果客户端已经在关闭过程中，返回 ErrShutdown
		return ErrShutdown
	}
	client.closing = true    // 标记客户端正在关闭
	return client.cc.Close() // 关闭底层的编解码器连接
}

// IsAvailable 检查客户端是否可用，即是否未关闭且未处于关闭过程中。
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing // 返回客户端是否处于可用状态
}

// registerCall 注册一个新的 RPC 调用，并为其分配一个唯一的序列号。
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown { // 如果客户端正在关闭或已关闭，返回 ErrShutdown
		return 0, ErrShutdown
	}
	call.Seq = client.seq           // 为调用分配当前序列号
	client.pending[call.Seq] = call // 将调用存储在 pending 中，以便后续匹配响应
	client.seq++                    // 序列号递增
	return call.Seq, nil            // 返回分配的序列号
}

// removeCall 根据序列号移除并返回对应的 Call 实例。
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq] // 根据序列号获取 Call 实例
	delete(client.pending, seq) // 从 pending 中移除该调用
	return call
}

// terminateCalls 终止所有未完成的调用，并将它们标记为指定的错误。
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true                // 标记客户端已完全关闭
	for _, call := range client.pending { // 遍历所有未完成的调用
		call.Error = err // 将调用的错误设置为传入的错误
		call.done()      // 通知调用完成
	}
}

// send 发送一个 RPC 请求到服务器。
func (client *Client) send(call *Call) {
	client.sending.Lock()
	defer client.sending.Unlock()

	// 注册调用并获取序列号
	seq, err := client.registerCall(call)
	if err != nil { // 如果注册失败（例如客户端已关闭），直接返回错误
		call.Error = err
		call.done()
		return
	}

	// 填充请求头部信息
	client.header.ServiceMethod = call.ServiceMethod // 设置服务方法名
	client.header.Seq = seq                          // 设置序列号
	client.header.Error = ""                         // 清空错误信息

	// 写入请求头部和参数
	if err := client.cc.Write(&client.header, call.Args); err != nil { // 将请求头部和参数写入连接
		call := client.removeCall(seq) // 移除调用
		if call != nil {
			call.Error = err // 设置错误信息
			call.done()      // 通知调用完成
		}
	}
}

// receive 是一个后台协程，用于接收服务器的响应。
// 它循环读取响应头部和正文，根据序列号匹配对应的 Call 实例，并通知调用完成。
// 如果发生错误，它会终止所有未完成的调用。
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil { // 从编解码器读取响应头部
			break
		}
		call := client.removeCall(h.Seq) // 根据序列号移除对应的 Call 实例
		switch {
		case call == nil: // 如果找不到对应的 Call，忽略响应正文
			err = client.cc.ReadBody(nil)
		case h.Error != "": // 如果响应头部包含错误信息
			call.Error = errors.New(h.Error) // 设置 Call 的错误信息
			err = client.cc.ReadBody(nil)    // 忽略响应正文
			call.done()                      // 通知调用完成
		default: // 正常响应
			err = client.cc.ReadBody(call.Reply) // 读取响应正文到 Call 的 Reply 中
			if err != nil {
				call.Error = errors.New("reading body " + err.Error()) // 设置读取错误
			}
			call.done() // 通知调用完成
		}
	}
	client.terminateCalls(err) // 如果发生错误，终止所有未完成的调用
}

// Go 发起一个异步的 RPC 调用，并返回一个 Call 实例。
// 它允许客户端在后台处理调用结果，而不会阻塞主线程。
// 如果 done 通道为 nil，会创建一个容量为 10 的缓冲通道；如果 done 通道是非缓冲的，会触发 panic。
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10) // 如果 done 通道为 nil，创建一个容量为 10 的缓冲通道
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered") // done 通道必须是缓冲的
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call) // 发送调用请求
	return call
}

// Call 发起一个同步的 RPC 调用，并返回调用结果的错误。
// 它使用上下文（context）来支持超时和取消操作。
// 如果上下文完成（超时或取消），调用会被终止；否则，等待调用完成并返回错误。
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1)) // 发起异步调用
	select {
	case <-ctx.Done(): // 如果上下文完成（超时或取消），取消调用
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done: // 等待调用完成
		return call.Error
	}
}

// parseOptions 解析客户端选项，确保选项的合法性。
// 如果未提供选项或选项为 nil，返回默认选项。
// 如果提供了多个选项，返回错误。
func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil // 如果未提供选项，返回默认选项
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1") // 只允许一个选项
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber // 设置默认魔数
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType // 设置默认编解码类型
	}
	return opt, nil
}

// NewClient 创建一个新的 RPC 客户端实例。
// 它根据提供的选项初始化客户端，并与服务器建立连接。
// 如果编解码类型无效或发送选项失败，返回错误。
func NewClient(conn net.Conn, opt *Option) (client *Client, err error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err = fmt.Errorf("invalid codec type %s", opt.CodecType) // 检查编解码类型是否有效
		log.Println("rpc client: codec error:", err)
		return
	}
	// 向服务器发送客户端选项
	if err = json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		return
	}
	return newClientCodec(f(conn), opt), nil
}

// newClientCodec 创建一个新的 Client 实例，并启动接收协程。
// 它初始化客户端的序列号、编解码器、选项和 pending 映射。
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,                      // 初始化序列号为 1
		cc:      cc,                     // 设置编解码器
		opt:     opt,                    // 设置客户端选项
		pending: make(map[uint64]*Call), // 初始化 pending 映射
	}
	go client.receive() // 启动接收协程
	return client
}

// clientResult 是一个辅助结构体，用于在 dialTimeout 函数中返回客户端连接的结果。
type clientResult struct {
	client *Client
	err    error
}

// newClientFunc 是一个函数类型，用于创建新的 Client 实例。
type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

// dialTimeout 是一个带超时的客户端连接函数。
// 它支持在指定时间内建立连接，并处理超时或连接失败的情况。
// 如果超时时间设置为 0，则不启用超时机制。
func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...) // 解析选项
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTimeout(network, address, opt.ConnectTimeout) // 带超时的网络连接
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = conn.Close() // 如果连接失败，关闭连接
		}
	}()
	ch := make(chan clientResult) // 创建一个通道用于接收连接结果
	go func() {
		client, err := f(conn, opt) // 调用连接函数
		ch <- clientResult{client: client, err: err}
	}()
	if opt.ConnectTimeout == 0 {
		result := <-ch // 如果没有超时设置，直接等待结果
		return result.client, result.err
	}
	select {
	case <-time.After(opt.ConnectTimeout): // 如果超时，返回错误
		return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
	case result := <-ch: // 等待连接结果
		return result.client, result.err
	}
}

// Dial 是一个高级接口，用于连接到指定的 RPC 服务器，并返回一个客户端实例。
// 它支持超时和默认选项。
// 它通过调用 dialTimeout 函数实现连接逻辑。
func Dial(network, address string, opts ...*Option) (*Client, error) {
	return dialTimeout(NewClient, network, address, opts...) // 调用 dialTimeout 函数
}
