// geerpc 包提供了一个简单的 RPC 服务器和客户端实现，支持超时机制和并发处理。
package geerpc

import (
	"context" // 提供上下文管理功能，用于超时和取消操作
	"net"     // 提供网络功能，用于建立服务器和客户端连接
	"os"
	"runtime"
	"strings" // 提供字符串操作功能，用于检查错误信息
	"testing" // 提供测试框架功能，用于编写测试用例
	"time"    // 提供时间相关功能，用于处理超时和延时操作
)

// Bar 是一个示例服务，包含一个 Timeout 方法，用于测试超时逻辑。
type Bar int

// Timeout 是 Bar 服务的一个方法，模拟一个耗时操作。
func (b Bar) Timeout(argv int, reply *int) error {
	time.Sleep(time.Second * 2) // 模拟耗时操作
	return nil
}

// startServer 启动一个 RPC 服务器，并将监听地址发送到 addr 通道。
func startServer(addr chan string) {
	var b Bar
	_ = Register(&b) // 注册 Bar 服务
	// 监听一个随机空闲的 TCP 端口
	l, _ := net.Listen("tcp", ":0")
	addr <- l.Addr().String() // 将监听地址发送到通道
	Accept(l)                 // 开始接受客户端连接
}

// TestClient_dialTimeout 测试客户端的 dialTimeout 函数，验证连接超时逻辑。
func TestClient_dialTimeout(t *testing.T) {
	t.Parallel()                    // 允许并行运行测试用例
	l, _ := net.Listen("tcp", ":0") // 监听一个随机端口

	// 定义一个模拟的连接函数，用于测试超时逻辑
	f := func(conn net.Conn, opt *Option) (client *Client, err error) {
		_ = conn.Close()            // 关闭连接
		time.Sleep(time.Second * 2) // 模拟耗时操作
		return nil, nil
	}

	// 测试连接超时
	t.Run("timeout", func(t *testing.T) {
		_, err := dialTimeout(f, "tcp", l.Addr().String(), &Option{ConnectTimeout: time.Second})
		_assert(err != nil && strings.Contains(err.Error(), "connect timeout"), "expect a timeout error")
	})

	// 测试连接超时时间为 0（无超时限制）
	t.Run("0", func(t *testing.T) {
		_, err := dialTimeout(f, "tcp", l.Addr().String(), &Option{ConnectTimeout: 0})
		_assert(err == nil, "0 means no limit")
	})
}

// TestClient_Call 测试客户端的 Call 方法，验证超时逻辑。
func TestClient_Call(t *testing.T) {
	t.Parallel()                // 允许并行运行测试用例
	addrCh := make(chan string) // 创建一个通道用于接收服务器地址
	go startServer(addrCh)      // 启动服务器
	addr := <-addrCh            // 获取服务器地址
	time.Sleep(time.Second)     // 等待服务器启动

	// 测试客户端超时
	t.Run("client timeout", func(t *testing.T) {
		client, _ := Dial("tcp", addr)                                   // 连接到服务器
		ctx, _ := context.WithTimeout(context.Background(), time.Second) // 设置超时上下文
		var reply int
		err := client.Call(ctx, "Bar.Timeout", 1, &reply) // 调用 Bar.Timeout 方法
		_assert(err != nil && strings.Contains(err.Error(), ctx.Err().Error()), "expect a timeout error")
	})

	// 测试服务器处理超时
	t.Run("server handle timeout", func(t *testing.T) {
		client, _ := Dial("tcp", addr, &Option{
			HandleTimeout: time.Second, // 设置服务器处理超时时间为 1 秒
		})
		var reply int
		err := client.Call(context.Background(), "Bar.Timeout", 1, &reply) // 调用 Bar.Timeout 方法
		_assert(err != nil && strings.Contains(err.Error(), "handle timeout"), "expect a timeout error")
	})
}

// TestXDial 测试 XDial 函数，验证其是否能正确连接到不同协议的 RPC 服务器。
func TestXDial(t *testing.T) {
	// 仅在 Linux 系统上运行测试，因为 Unix 套接字在非 Linux 系统上可能不可用
	if runtime.GOOS == "linux" {
		ch := make(chan struct{})  // 创建一个通道用于同步服务器启动
		addr := "/tmp/geerpc.sock" // 定义 Unix 套接字的地址

		// 启动一个 Unix 套接字服务器
		go func() {
			_ = os.Remove(addr)                // 如果文件已存在，删除它
			l, err := net.Listen("unix", addr) // 监听 Unix 套接字
			if err != nil {
				t.Fatal("failed to listen unix socket") // 如果监听失败，记录错误
			}
			ch <- struct{}{} // 通知测试函数服务器已启动
			Accept(l)        // 开始接受客户端连接
		}()

		// 等待服务器启动
		<-ch

		// 使用 XDial 连接到 Unix 套接字服务器
		_, err := XDial("unix@" + addr)
		_assert(err == nil, "failed to connect unix socket") // 验证连接是否成功
	}
}
