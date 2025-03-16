package log

import (
	"io/ioutil" // 提供对文件和文件系统的I/O操作
	"log"       // 提供简单的日志功能
	"os"        // 提供操作系统相关的功能
	"sync"      // 提供基本的同步原语，如互斥锁
)

// 定义全局日志对象
var (
	// errorLog 是用于记录错误日志的Logger，输出到标准输出，并带有红色的"[error]"前缀
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)

	// infoLog 是用于记录信息日志的Logger，输出到标准输出，并带有蓝色的"[info ]"前缀
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)

	// loggers 是一个包含所有日志对象的切片，方便统一管理
	loggers = []*log.Logger{errorLog, infoLog}

	// mu 是一个互斥锁，用于在设置日志级别时同步对日志输出的修改
	mu sync.Mutex
)

// 定义日志输出的快捷方法
var (
	// Error 是一个快捷方法，用于输出错误日志
	Error = errorLog.Println

	// Errorf 是一个快捷方法，用于格式化输出错误日志
	Errorf = errorLog.Printf

	// Info 是一个快捷方法，用于输出信息日志
	Info = infoLog.Println

	// Infof 是一个快捷方法，用于格式化输出信息日志
	Infof = infoLog.Printf
)

// 定义日志级别常量
const (
	// InfoLevel 表示信息级别日志
	InfoLevel = iota

	// ErrorLevel 表示错误级别日志
	ErrorLevel

	// Disabled 表示禁用所有日志输出
	Disabled
)

// SetLevel 用于设置日志级别
func SetLevel(level int) {
	mu.Lock() // 加锁，确保在设置日志级别时线程安全
	defer mu.Unlock()

	// 默认情况下，所有日志都输出到标准输出
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	// 根据设置的日志级别调整日志输出
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard) // 如果错误日志级别被禁用，将其输出丢弃
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard) // 如果信息日志级别被禁用，将其输出丢弃
	}
}
