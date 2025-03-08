package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// trace 函数用于生成带有调用栈信息的错误消息。
// 它通过 runtime 包获取调用栈信息，并将其格式化为字符串。
func trace(message string) string {
	// pcs 是一个数组，用于存储调用栈中的程序计数器（PC）值。
	// [32]uintptr 表示数组的大小为 32，每个元素是一个 uintptr 类型的值。
	var pcs [32]uintptr
	// runtime.Callers 获取调用栈信息，跳过前 3 个调用者。
	// 这里跳过 3 个调用者的原因是：
	// 1. trace 函数本身
	// 2. Recovery 函数中的 defer 调用
	// 3. Recovery 函数本身
	n := runtime.Callers(3, pcs[:])

	// 使用 strings.Builder 动态构建字符串，提高性能。
	var str strings.Builder
	// 将错误消息和固定的字符串 "Traceback:" 写入到 str 中。
	str.WriteString(message + "\nTraceback:")

	// 遍历 pcs 数组中存储的调用栈信息（只遍历实际获取的数量 n）。
	for _, pc := range pcs[:n] {
		// runtime.FuncForPC 根据程序计数器 pc 获取对应的函数信息。
		fn := runtime.FuncForPC(pc)
		// fn.FileLine 获取该函数所在的文件路径和行号。
		file, line := fn.FileLine(pc)
		// 使用 fmt.Sprintf 将文件路径和行号格式化为字符串，并写入到 str 中。
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	// 返回构建好的字符串，包含错误消息和调用栈信息。
	return str.String()
}

// Recovery 是一个中间件函数，用于捕获和处理运行时的 panic。
// 它通过 defer 和 recover 机制捕获 panic，并返回 500 内部服务器错误。
func Recovery() HandlerFunc {
	return func(c *Context) {
		// 使用 defer 延迟执行，确保在函数返回时捕获 panic。
		defer func() {
			// 如果发生 panic，recover 会捕获 panic 的错误信息。
			if err := recover(); err != nil {
				// 将 panic 的错误信息格式化为字符串。
				message := fmt.Sprintf("%s", err)
				// 使用 trace 函数生成带有调用栈信息的错误消息。
				log.Printf("%s\n\n", trace(message))
				// 调用 c.Fail 返回 500 内部服务器错误，并向客户端发送错误信息。
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		// 调用 c.Next() 继续执行后续的中间件或处理器。
		c.Next()
	}
}
