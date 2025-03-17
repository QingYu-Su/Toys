package xclient

import (
	"context"
	. "geerpc" // 引入 geerpc 包，方便直接使用其类型和方法
	"io"       // 提供基本的 I/O 操作接口
	"reflect"  // 提供反射功能，用于动态操作类型和值
	"sync"     // 提供同步原语，用于并发控制
)

// XClient 是一个扩展的 RPC 客户端，支持服务发现和多种选择模式。
type XClient struct {
	d       Discovery          // 服务发现接口，用于动态获取服务器地址
	mode    SelectMode         // 选择模式（随机选择或轮询选择）
	opt     *Option            // RPC 客户端的配置选项
	mu      sync.Mutex         // 保护以下字段的互斥锁
	clients map[string]*Client // 存储已连接的 RPC 客户端实例
}

// 确保 XClient 实现了 io.Closer 接口，以便可以关闭所有客户端连接。
var _ io.Closer = (*XClient)(nil)

// NewXClient 创建一个 XClient 实例，用于管理 RPC 客户端和服务发现。
func NewXClient(d Discovery, mode SelectMode, opt *Option) *XClient {
	return &XClient{
		d:       d,                        // 设置服务发现接口
		mode:    mode,                     // 设置选择模式
		opt:     opt,                      // 设置配置选项
		clients: make(map[string]*Client), // 初始化客户端映射
	}
}

// Close 关闭所有已连接的 RPC 客户端。
func (xc *XClient) Close() error {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	for key, client := range xc.clients {
		// 关闭每个客户端连接，忽略关闭时的错误
		_ = client.Close()
		delete(xc.clients, key) // 从映射中删除客户端
	}
	return nil
}

// dial 建立与指定 RPC 服务器的连接。
// 如果已存在可用的客户端实例，则直接返回；否则重新建立连接。
func (xc *XClient) dial(rpcAddr string) (*Client, error) {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	client, ok := xc.clients[rpcAddr] // 检查是否已存在客户端实例
	if ok && !client.IsAvailable() {  // 如果客户端不可用
		_ = client.Close()          // 关闭不可用的客户端
		delete(xc.clients, rpcAddr) // 从映射中删除
		client = nil
	}
	if client == nil { // 如果没有可用的客户端实例
		var err error
		client, err = XDial(rpcAddr, xc.opt) // 使用 XDial 建立连接
		if err != nil {
			return nil, err
		}
		xc.clients[rpcAddr] = client // 将新客户端存储到映射中
	}
	return client, nil
}

// call 调用指定的 RPC 服务器的指定方法。
// 它通过 dial 方法获取一个可用的 RPC 客户端，然后调用该客户端的 Call 方法。
func (xc *XClient) call(rpcAddr string, ctx context.Context, serviceMethod string, args, reply interface{}) error {
	client, err := xc.dial(rpcAddr) // 获取一个可用的 RPC 客户端
	if err != nil {
		return err
	}
	return client.Call(ctx, serviceMethod, args, reply) // 调用指定的 RPC 方法
}

// Call 调用指定的服务方法，并等待调用完成。
// 它根据选择模式从服务发现中获取一个服务器地址，然后调用该服务器的指定方法。
func (xc *XClient) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	rpcAddr, err := xc.d.Get(xc.mode) // 根据选择模式获取一个服务器地址
	if err != nil {
		return err
	}
	return xc.call(rpcAddr, ctx, serviceMethod, args, reply) // 调用指定的 RPC 方法
}

// Broadcast 调用注册在服务发现中的所有服务器的指定方法。
// 它并发调用每个服务器的指定方法，并返回第一个成功的结果。
// 如果所有调用都失败，则返回最后一个错误。
func (xc *XClient) Broadcast(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	servers, err := xc.d.GetAll() // 获取所有服务器地址
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex // 保护 e 和 replyDone
	var e error
	replyDone := reply == nil              // 如果 reply 为 nil，则不需要设置值
	ctx, cancel := context.WithCancel(ctx) // 创建一个可取消的上下文
	for _, rpcAddr := range servers {
		wg.Add(1)
		go func(rpcAddr string) {
			defer wg.Done()
			var clonedReply interface{}
			if reply != nil {
				// 如果 reply 不为 nil，则克隆 reply 的值
				clonedReply = reflect.New(reflect.ValueOf(reply).Elem().Type()).Interface()
			}
			err := xc.call(rpcAddr, ctx, serviceMethod, args, clonedReply) // 调用指定的 RPC 方法
			mu.Lock()
			if err != nil && e == nil {
				// 如果当前调用失败且尚未记录错误，则记录错误并取消其他调用
				e = err
				cancel()
			}
			if err == nil && !replyDone {
				// 如果当前调用成功且 replyDone 为 false，则设置 reply 的值
				reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(clonedReply).Elem())
				replyDone = true
			}
			mu.Unlock()
		}(rpcAddr)
	}
	wg.Wait() // 等待所有调用完成
	return e  // 返回第一个错误或 nil
}
