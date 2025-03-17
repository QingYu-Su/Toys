package geerpc

import (
	"encoding/json" // 用于 JSON 编解码，主要在客户端和服务端之间传递配置选项
	"errors"        // 提供错误处理功能
	"fmt"           // 用于格式化输出
	"geerpc/codec"  // 引入自定义的 codec 包，用于处理 RPC 请求和响应的编解码
	"io"            // 提供基本的 I/O 操作接口，确保 Client 实现 io.Closer
	"log"           // 用于记录日志信息，方便调试和运行时监控
	"net"           // 提供网络功能，用于建立和管理客户端与服务端的连接
	"sync"          // 提供同步原语，用于处理并发请求和响应
)

// Call 表示一次 RPC 调用的上下文，包含请求和响应的所有信息。
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
	opt      *Option          // 客户端配置选项，例如编解码类型
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

// receive 接收服务器的响应，并处理响应内容。
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header                              // 响应头部
		if err = client.cc.ReadHeader(&h); err != nil { // 读取响应头部
			break
		}
		call := client.removeCall(h.Seq) // 根据序列号获取对应的调用
		switch {
		case call == nil: // 如果找不到对应的调用，忽略响应正文
			err = client.cc.ReadBody(nil)
		case h.Error != "": // 如果响应头部包含错误信息
			call.Error = errors.New(h.Error) // 设置错误信息
			err = client.cc.ReadBody(nil)    // 读取响应正文（忽略内容）
			call.done()                      // 通知调用完成
		default: // 正常响应
			err = client.cc.ReadBody(call.Reply) // 读取响应正文到 Reply
			if err != nil {
				call.Error = errors.New("reading body " + err.Error()) // 设置错误信息
			}
			call.done() // 通知调用完成
		}
	}
	client.terminateCalls(err) // 如果发生错误，终止所有未完成的调用
}

// Go 发起一个异步的 RPC 调用，并返回一个 Call 实例。
// 这个方法是非阻塞的，它会立即返回一个 Call 对象，而不会等待 RPC 调用完成。
// 调用结果（成功或失败）会通过 Call 的 Done 通道通知调用者。
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	// 如果 done 通道为 nil，则创建一个缓冲通道，用于接收调用完成的通知。
	// 缓冲区大小设置为 10，这意味着可以同时处理 10 个未完成的调用。
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		// 如果 done 通道是非缓冲的（即容量为 0），则触发 panic。
		// 这是因为非缓冲通道会导致发送方阻塞，直到有接收方准备接收。
		// 在异步调用场景中，这可能会导致死锁，因此强制要求 done 通道必须是缓冲的。
		log.Panic("rpc client: done channel is unbuffered")
	}

	// 创建一个新的 Call 实例，包含调用的所有必要信息：
	// - ServiceMethod: 要调用的服务方法名（格式为 "Service.Method"）。
	// - Args: 请求参数。
	// - Reply: 用于存储响应结果的变量。
	// - Done: 用于通知调用完成的通道。
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}

	// 调用 send 方法将 Call 实例发送到服务器。
	// send 方法会负责序列化请求数据，并将其写入底层连接。
	client.send(call)

	// 返回 Call 实例，调用者可以通过 Call.Done 通道接收调用完成的通知，
	// 并通过 Call.Error 检查调用是否成功，通过 Call.Reply 获取响应结果。
	return call
}

// Call 发起一个同步的 RPC 调用，并返回调用结果的错误。
// 这是一个阻塞方法，会等待 RPC 调用完成后再返回。
func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}

// parseOptions 解析客户端选项，确保选项的合法性。
// 如果未提供选项或选项为 nil，则使用默认选项。
func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// NewClient 创建一个新的 RPC 客户端实例。
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	// send options with server
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

// newClientCodec 创建一个新的 Client 实例，并启动接收协程。
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,                      // 初始化序列号为 1
		cc:      cc,                     // 设置编解码器
		opt:     opt,                    // 设置客户端选项
		pending: make(map[uint64]*Call), // 初始化 pending 映射
	}
	go client.receive() // 启动接收协程，用于处理服务器的响应
	return client
}

// Dial 连接到指定的 RPC 服务器，并返回一个客户端实例。
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// close the connection if client is nil
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	return NewClient(conn, opt)
}
