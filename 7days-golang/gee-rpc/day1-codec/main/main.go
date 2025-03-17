package main

import (
	"encoding/json" // 用于 JSON 编解码
	"fmt"           // 用于格式化输出
	"geerpc"        // 引入自定义的 geerpc 包，提供 RPC 服务器和客户端功能
	"geerpc/codec"  // 引入自定义的 codec 包，提供编解码器功能
	"log"           // 用于记录日志信息
	"net"           // 提供网络相关功能
	"time"          // 提供时间相关功能
)

// startServer 启动一个 RPC 服务器，并将监听地址发送到 addr 通道
func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0") // 随机选择一个空闲端口进行监听
	if err != nil {
		log.Fatal("network error:", err) // 如果监听失败，记录错误并退出
	}
	log.Println("start rpc server on", l.Addr()) // 记录服务器启动信息
	addr <- l.Addr().String()                    // 将监听地址发送到通道
	geerpc.Accept(l)                             // 调用 geerpc.Accept 开始接受客户端连接
}

func main() {
	addr := make(chan string) // 创建一个通道用于接收服务器的监听地址
	go startServer(addr)      // 启动一个 RPC 服务器

	// in fact, following code is like a simple geerpc client
	conn, _ := net.Dial("tcp", <-addr)  // 连接到服务器
	defer func() { _ = conn.Close() }() // 确保连接在函数退出时关闭

	time.Sleep(time.Second) // 等待服务器启动完成
	// send options
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption) // 向服务器发送默认的 RPC 选项
	cc := codec.NewGobCodec(conn)                          // 创建一个 Gob 编解码器

	// send request & receive response
	for i := 0; i < 5; i++ { // 发送 5 个请求并接收响应
		h := &codec.Header{
			ServiceMethod: "Foo.Sum", // 指定服务方法名
			Seq:           uint64(i), // 设置请求序列号
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq)) // 写入请求头部和正文
		_ = cc.ReadHeader(h)                                 // 读取响应头部
		var reply string
		_ = cc.ReadBody(&reply)      // 读取响应正文
		log.Println("reply:", reply) // 记录响应内容
	}
}
