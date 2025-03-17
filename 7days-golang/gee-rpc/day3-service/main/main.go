package main

import (
	"geerpc" // 引入自定义的 geerpc 包，提供 RPC 服务和客户端功能
	"log"    // 用于记录日志信息
	"net"    // 提供网络功能，用于启动服务器和建立客户端连接
	"sync"   // 提供同步原语，用于管理并发请求
	"time"   // 提供时间相关功能，用于延时等待
)

// Foo 是一个示例服务，包含一个符合 RPC 方法签名的导出方法 Sum。
type Foo int

// Args 是一个结构体，表示 RPC 方法的请求参数。
type Args struct{ Num1, Num2 int }

// Sum 是 Foo 服务的 RPC 方法，计算两个整数的和并将结果存储在 reply 中。
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2 // 计算结果并存储到 reply 中
	return nil                     // 返回 nil 表示调用成功
}

// startServer 启动一个 RPC 服务器，并将监听地址发送到 addr 通道。
func startServer(addr chan string) {
	var foo Foo                                   // 创建一个 Foo 类型的服务实例
	if err := geerpc.Register(&foo); err != nil { // 将服务注册到默认服务器
		log.Fatal("register error:", err)
	}
	// 监听一个随机空闲的 TCP 端口
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr()) // 记录服务器启动的地址
	addr <- l.Addr().String()                    // 将监听地址发送到 addr 通道
	geerpc.Accept(l)                             // 开始接受客户端连接
}

func main() {
	log.SetFlags(0)           // 设置日志格式，去掉时间戳等默认前缀
	addr := make(chan string) // 创建一个通道用于接收服务器的监听地址
	go startServer(addr)      // 启动一个 RPC 服务器

	// 连接到服务器
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }() // 确保程序退出时关闭客户端连接

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
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err) // 如果调用失败，记录错误并退出
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply) // 记录调用结果
		}(i)
	}
	wg.Wait() // 等待所有并发请求完成
}
