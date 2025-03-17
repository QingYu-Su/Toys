package main

import (
	"fmt"    // 用于格式化输出
	"geerpc" // 引入自定义的 geerpc 包，提供 RPC 客户端和服务器功能
	"log"    // 用于记录日志信息
	"net"    // 提供网络功能，用于启动服务器和建立客户端连接
	"sync"   // 提供同步原语，用于管理并发请求
	"time"   // 提供时间相关功能，用于延时等待
)

// startServer 启动一个 RPC 服务器，并将监听地址发送到 addr 通道。
// 该函数会监听一个随机空闲端口，并通过 geerpc.Accept 开始接受客户端连接。
func startServer(addr chan string) {
	// 监听一个随机空闲的 TCP 端口
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err) // 如果监听失败，记录错误并退出
	}
	log.Println("start rpc server on", l.Addr()) // 记录服务器启动的地址
	addr <- l.Addr().String()                    // 将监听地址发送到 addr 通道
	geerpc.Accept(l)                             // 调用 geerpc.Accept 开始接受客户端连接
}

func main() {
	log.SetFlags(0) // 设置日志格式，去掉时间戳等默认前缀

	// 创建一个通道用于接收服务器的监听地址
	addr := make(chan string)
	// 启动一个 RPC 服务器，并将监听地址发送到 addr 通道
	go startServer(addr)

	// 连接到服务器
	client, _ := geerpc.Dial("tcp", <-addr) // 从 addr 通道接收服务器地址并建立连接
	defer func() { _ = client.Close() }()   // 确保程序退出时关闭客户端连接

	// 等待 1 秒，确保服务器已经启动
	time.Sleep(time.Second)

	// 使用 WaitGroup 管理并发请求
	var wg sync.WaitGroup
	// 发送 5 个并发请求并接收响应
	for i := 0; i < 5; i++ {
		wg.Add(1) // 增加 WaitGroup 的计数
		go func(i int) {
			defer wg.Done() // 请求完成时减少 WaitGroup 的计数
			// 构造请求参数
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string // 用于存储响应结果
			// 调用服务器的 Foo.Sum 方法
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err) // 如果调用失败，记录错误并退出
			}
			log.Println("reply:", reply) // 记录响应结果
		}(i)
	}
	wg.Wait() // 等待所有并发请求完成
}
