package main

import (
	"context"         // 提供上下文管理功能，用于超时和取消操作
	"geerpc"          // 提供 RPC 服务和客户端功能
	"geerpc/registry" // 提供注册中心功能
	"geerpc/xclient"  // 提供扩展的 RPC 客户端功能
	"log"             // 提供日志记录功能
	"net"             // 提供网络功能，用于启动服务器
	"net/http"        // 提供 HTTP 服务器功能
	"sync"            // 提供同步原语，用于并发控制
	"time"            // 提供时间相关功能
)

// Foo 是一个示例服务，包含两个方法：Sum 和 Sleep。
type Foo int

// Args 是一个结构体，表示 RPC 方法的请求参数。
type Args struct{ Num1, Num2 int }

// Sum 是 Foo 服务的一个方法，计算两个整数的和并将结果存储在 reply 中。
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2 // 计算结果并存储到 reply 中
	return nil                     // 返回 nil 表示调用成功
}

// Sleep 是 Foo 服务的一个方法，模拟一个耗时操作。
func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Duration(args.Num1) * time.Second) // 模拟耗时操作
	*reply = args.Num1 + args.Num2                     // 计算结果并存储到 reply 中
	return nil                                         // 返回 nil 表示调用成功
}

// startRegistry 启动一个注册中心，监听指定的地址。
func startRegistry(wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":9999") // 监听 TCP 端口 9999
	registry.HandleHTTP()              // 注册 HTTP 处理器
	wg.Done()                          // 通知 WaitGroup 当前协程已启动
	_ = http.Serve(l, nil)             // 启动 HTTP 服务器
}

// startServer 启动一个 RPC 服务器，注册 Foo 服务，并向注册中心发送心跳。
func startServer(registryAddr string, wg *sync.WaitGroup) {
	var foo Foo                                                   // 创建一个 Foo 类型的服务实例
	l, _ := net.Listen("tcp", ":0")                               // 监听一个随机空闲的 TCP 端口
	server := geerpc.NewServer()                                  // 创建一个新的 RPC 服务器实例
	_ = server.Register(&foo)                                     // 将 Foo 服务注册到服务器
	registry.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0) // 向注册中心发送心跳
	wg.Done()                                                     // 通知 WaitGroup 当前协程已启动
	server.Accept(l)                                              // 开始接受客户端连接
}

// foo 是一个辅助函数，用于调用 RPC 方法。
func foo(xc *xclient.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int // 用于存储响应结果
	var err error // 用于存储调用过程中发生的错误
	switch typ {
	case "call": // 单次调用
		err = xc.Call(ctx, serviceMethod, args, &reply) // 调用指定的服务方法
	case "broadcast": // 广播调用
		err = xc.Broadcast(ctx, serviceMethod, args, &reply) // 广播调用所有服务器
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err) // 如果调用失败，记录错误
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply) // 如果调用成功，记录结果
	}
}

// call 测试单次调用 Foo.Sum 方法。
func call(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)      // 创建一个 GeeRegistryDiscovery 实例
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil) // 创建一个 XClient 实例
	defer func() { _ = xc.Close() }()                      // 确保关闭 XClient
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "call", "Foo.Sum", &Args{Num1: i, Num2: i * i}) // 单次调用 Foo.Sum 方法
		}(i)
	}
	wg.Wait()
}

// broadcast 测试广播调用 Foo.Sum 和 Foo.Sleep 方法。
func broadcast(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)      // 创建一个 GeeRegistryDiscovery 实例
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil) // 创建一个 XClient 实例
	defer func() { _ = xc.Close() }()                      // 确保关闭 XClient
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "broadcast", "Foo.Sum", &Args{Num1: i, Num2: i * i}) // 广播调用 Foo.Sum 方法
			// expect 2 - 5 timeout
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)  // 设置超时时间为 2 秒
			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{Num1: i, Num2: i * i}) // 广播调用 Foo.Sleep 方法
		}(i)
	}
	wg.Wait()
}

func main() {
	log.SetFlags(0)                                           // 设置日志格式，去掉时间戳等默认前缀
	registryAddr := "http://localhost:9999/_geerpc_/registry" // 注册中心地址
	var wg sync.WaitGroup

	wg.Add(1)
	go startRegistry(&wg) // 启动注册中心
	wg.Wait()             // 等待注册中心启动

	time.Sleep(time.Second) // 等待注册中心启动完成
	wg.Add(2)
	go startServer(registryAddr, &wg) // 启动第一个 RPC 服务器
	go startServer(registryAddr, &wg) // 启动第二个 RPC 服务器
	wg.Wait()                         // 等待服务器启动完成

	time.Sleep(time.Second) // 等待服务器启动完成
	call(registryAddr)      // 测试单次调用
	broadcast(registryAddr) // 测试广播调用
}
