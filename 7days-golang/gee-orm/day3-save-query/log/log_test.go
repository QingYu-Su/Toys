package log

import (
	"os"      // 提供操作系统相关的功能
	"testing" // 提供单元测试功能
)

// TestSetLevel 是一个单元测试函数，用于测试 SetLevel 函数的功能。
func TestSetLevel(t *testing.T) {
	// 设置日志级别为 ErrorLevel
	SetLevel(ErrorLevel)
	// 检查设置是否正确：
	// - infoLog 应该被丢弃（不输出到标准输出）
	// - errorLog 应该仍然输出到标准输出
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level") // 如果检查失败，终止测试并报错
	}

	// 设置日志级别为 Disabled
	SetLevel(Disabled)
	// 检查设置是否正确：
	// - infoLog 和 errorLog 都应该被丢弃（不输出到标准输出）
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level") // 如果检查失败，终止测试并报错
	}
}
