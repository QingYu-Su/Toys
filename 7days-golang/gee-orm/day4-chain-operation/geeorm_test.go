package geeorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3" // 引入 SQLite 驱动，用于测试 SQLite 数据库
)

// OpenDB 是一个辅助函数，用于打开一个 SQLite 数据库连接，并返回一个 Engine 实例。
// 它会在测试失败时调用 t.Fatal，因此可以确保测试环境的正确性。
func OpenDB(t *testing.T) *Engine {
	t.Helper() // 标记为辅助函数，便于调试和日志记录
	// 创建一个新的 Engine 实例，连接到 SQLite 数据库。
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err) // 如果连接失败，终止测试并报错
	}
	return engine
}

// TestNewEngine 测试 NewEngine 函数的功能。
// 它检查是否能够成功创建一个 Engine 实例并连接到数据库。
func TestNewEngine(t *testing.T) {
	// 使用 OpenDB 打开一个数据库连接。
	engine := OpenDB(t)
	// 在测试完成后关闭数据库连接。
	defer engine.Close()
	// 如果代码执行到这里，说明 NewEngine 已成功创建 Engine 实例。
}
