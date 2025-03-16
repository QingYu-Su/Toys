package geeorm

import (
	"errors"
	"geeorm/session"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3" // 引入 SQLite 驱动，用于测试 SQLite 数据库
)

// OpenDB 打开一个 SQLite 数据库连接，并返回一个 Engine 实例。
func OpenDB(t *testing.T) *Engine {
	t.Helper()                                    // 标记为辅助函数
	engine, err := NewEngine("sqlite3", "gee.db") // 创建一个新的 Engine 实例
	if err != nil {
		t.Fatal("failed to connect", err) // 如果连接失败，终止测试
	}
	return engine
}

// User 是一个测试用的结构体，模拟数据库中的用户表。
type User struct {
	Name string `geeorm:"PRIMARY KEY"` // 主键字段
	Age  int    // 年龄字段
}

// TestEngine_Transaction 测试 Engine 的 Transaction 方法。
func TestEngine_Transaction(t *testing.T) {
	// 使用 t.Run 分别测试事务的回滚和提交
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

// transactionRollback 测试事务回滚是否正确执行。
func transactionRollback(t *testing.T) {
	engine := OpenDB(t)              // 打开数据库连接
	defer engine.Close()             // 测试结束后关闭连接
	s := engine.NewSession()         // 创建一个新的 Session 实例
	_ = s.Model(&User{}).DropTable() // 删除 User 表（如果存在）

	// 执行事务函数，故意制造一个错误以触发回滚
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()  // 创建 User 表
		_, err = s.Insert(&User{"Tom", 18}) // 插入一条记录
		return nil, errors.New("Error")     // 返回错误，触发回滚
	})

	// 检查事务是否正确回滚
	if err == nil || s.HasTable() { // 如果没有错误或表仍然存在，说明回滚失败
		t.Fatal("failed to rollback")
	}
}

// transactionCommit 测试事务提交是否正确执行。
func transactionCommit(t *testing.T) {
	engine := OpenDB(t)              // 打开数据库连接
	defer engine.Close()             // 测试结束后关闭连接
	s := engine.NewSession()         // 创建一个新的 Session 实例
	_ = s.Model(&User{}).DropTable() // 删除 User 表（如果存在）

	// 执行事务函数
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()  // 创建 User 表
		_, err = s.Insert(&User{"Tom", 18}) // 插入一条记录
		return                              // 正常结束事务，触发提交
	})

	// 检查事务是否正确提交
	u := &User{}                       // 创建一个 User 实例用于存储查询结果
	_ = s.First(u)                     // 查询第一条记录
	if err != nil || u.Name != "Tom" { // 如果查询失败或记录不正确，说明提交失败
		t.Fatal("failed to commit")
	}
}

// TestEngine_Migrate 测试 Engine 的 Migrate 方法是否能够正确迁移表结构。
func TestEngine_Migrate(t *testing.T) {
	// 打开数据库连接
	engine := OpenDB(t)
	defer engine.Close() // 确保在测试结束时关闭数据库连接

	// 创建一个新的 Session 实例
	s := engine.NewSession()

	// 初始化测试环境：删除 User 表（如果存在），并创建一个包含额外字段的 User 表
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	_, _ = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()

	// 调用 Migrate 方法，传入 User 结构体的实例
	// Migrate 方法会根据 User 结构体的定义调整数据库表结构
	engine.Migrate(&User{})

	// 查询 User 表的列名
	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	columns, _ := rows.Columns()

	// 检查表结构是否正确迁移
	// User 结构体定义了 Name 和 Age 两个字段，因此期望的列名应该是 ["Name", "Age"]
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", columns) // 如果列名不匹配，测试失败
	}
}
