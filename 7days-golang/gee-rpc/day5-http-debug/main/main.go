package main

import (
	"context"  // 提供上下文管理功能，用于超时和取消操作
	"geerpc"   // 提供 RPC 客户端和服务器功能
	"log"      // 提供日志记录功能
	"net"      // 提供网络功能，用于建立服务器和客户端连接
	"net/http" // 提供 HTTP 服务器和客户端功能
	"sync"     // 提供同步原语，用于管理并发请求
	"time"     // 提供时间相关功能，用于处理延时操作
)

// Foo 是一个示例服务，包含一个 Sum 方法，用于计算两个整数的和。
type Foo int

// Args 是一个结构体，表示 RPC 方法的请求参数。
type Args struct{ Num1, Num2 int }

// Sum 是 Foo 服务的一个方法，计算两个整数的和并将结果存储在 reply 中。
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2 // 计算结果并存储到 reply 中
	return nil                     // 返回 nil 表示调用成功
}

// startServer 启动一个 HTTP RPC 服务器，并将监听地址发送到 addrCh 通道。
func startServer(addrCh chan string) {
	var foo Foo                        // 创建一个 Foo 类型的服务实例
	l, _ := net.Listen("tcp", ":9999") // 监听 TCP 端口 9999
	_ = geerpc.Register(&foo)          // 将 Foo 服务注册到默认服务器
	geerpc.HandleHTTP()                // 注册 HTTP RPC 处理器
	addrCh <- l.Addr().String()        // 将监听地址发送到通道
	_ = http.Serve(l, nil)             // 启动 HTTP 服务器
}

// call 连接到 HTTP RPC 服务器，并并发调用 Foo.Sum 方法。
func call(addrCh chan string) {
	client, _ := geerpc.DialHTTP("tcp", <-addrCh) // 连接到 HTTP RPC 服务器
	defer func() { _ = client.Close() }()         // 确保程序退出时关闭客户端连接

	time.Sleep(time.Second) // 等待 1 秒，确保服务器已经启动

	// 发送请求并接收响应
	var wg sync.WaitGroup // 使用 WaitGroup 管理并发请求
	for i := 0; i < 5; i++ {
		wg.Add(1) // 增加 WaitGroup 的计数
		go func(i int) {
			defer wg.Done()                     // 请求完成时减少 WaitGroup 的计数
			args := &Args{Num1: i, Num2: i * i} // 构造请求参数
			var reply int                       // 用于存储响应结果
			// 调用 Foo.Sum 方法
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err) // 如果调用失败，记录错误并退出
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply) // 记录调用结果
		}(i)
	}
	wg.Wait() // 等待所有并发请求完成
}

func main() {
	log.SetFlags(0)         // 设置日志格式，去掉时间戳等默认前缀
	ch := make(chan string) // 创建一个通道用于接收服务器的监听地址
	go call(ch)             // 启动客户端
	startServer(ch)         // 启动服务器
}
