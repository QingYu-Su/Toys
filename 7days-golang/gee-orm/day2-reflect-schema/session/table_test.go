package session

import "testing"

// User 是一个测试用的结构体，用于模拟数据库中的表。
// 它包含两个字段：Name 和 Age，其中 Name 字段标记为 PRIMARY KEY。
type User struct {
	Name string `geeorm:"PRIMARY KEY"` // Name 字段，标记为表的主键
	Age  int    // Age 字段
}

// TestSession_CreateTable 测试 Session 的 CreateTable 方法。
func TestSession_CreateTable(t *testing.T) {
	// 创建一个新的 Session 实例，并设置模型为 User 结构体。
	s := NewSession().Model(&User{})

	// 如果表存在，则先删除表，确保测试环境干净。
	_ = s.DropTable()

	// 创建表。
	_ = s.CreateTable()

	// 检查表是否创建成功。
	if !s.HasTable() {
		t.Fatal("Failed to create table User") // 如果表未创建成功，终止测试并报错
	}
}
