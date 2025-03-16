package schema

import (
	"geeorm/dialect" // 引入方言包，用于获取数据库方言
	"testing"        // 引入测试包，用于编写单元测试
)

// User 是一个测试用的结构体，用于模拟数据库中的表。
// 它包含两个字段：Name 和 Age。
type User struct {
	Name string `geeorm:"PRIMARY KEY"` // Name 字段，标记为 PRIMARY KEY
	Age  int    // Age 字段
}

// TestDial 是一个全局变量，用于存储测试用的 SQLite3 方言。
var TestDial, _ = dialect.GetDialect("sqlite3")

// TestParse 测试 Parse 函数的功能。
func TestParse(t *testing.T) {
	// 使用 Parse 函数解析 User 结构体，并传入 SQLite3 方言。
	schema := Parse(&User{}, TestDial)

	// 检查解析后的 Schema 对象是否符合预期：
	// 1. 结构体名称应为 "User"。
	// 2. 字段数量应为 2（Name 和 Age）。
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct") // 如果不符合预期，终止测试并报错
	}

	// 检查 Name 字段的 Tag 是否为 "PRIMARY KEY"。
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key") // 如果不符合预期，终止测试并报错
	}
}
