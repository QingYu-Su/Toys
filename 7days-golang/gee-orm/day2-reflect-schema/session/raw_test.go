package session

import (
	"database/sql" // 提供数据库操作功能
	"os"           // 提供操作系统相关的功能
	"testing"      // 提供单元测试功能

	"geeorm/dialect" // 引入方言包，用于获取数据库方言

	_ "github.com/mattn/go-sqlite3" // 引入 SQLite 驱动，用于测试 SQLite 数据库
)

// TestDB 是一个全局变量，用于存储测试用的 SQLite 数据库连接。
var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3") // 获取 SQLite3 方言
)

// TestMain 是测试主函数，用于初始化和清理测试环境。
func TestMain(m *testing.M) {
	// 打开一个 SQLite 数据库连接，用于测试。
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	// 运行测试用例。
	code := m.Run()
	// 关闭数据库连接。
	_ = TestDB.Close()
	// 退出程序，返回测试结果代码。
	os.Exit(code)
}

// NewSession 创建一个新的 Session 实例，用于测试。
func NewSession() *Session {
	// 使用 TestDB 和 TestDial 创建一个新的 Session 实例。
	return New(TestDB, TestDial)
}

// TestSession_Exec 测试 Session 的 Exec 方法，用于执行 SQL 语句。
func TestSession_Exec(t *testing.T) {
	s := NewSession() // 创建一个新的 Session 实例
	// 删除 User 表（如果存在）。
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	// 创建 User 表。
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// 向 User 表插入两条数据。
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	// 检查插入的行数是否为 2。
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count) // 如果检查失败，终止测试并报错
	}
}

// TestSession_QueryRows 测试 Session 的 QueryRows 方法，用于查询多行数据。
func TestSession_QueryRows(t *testing.T) {
	s := NewSession() // 创建一个新的 Session 实例
	// 删除 User 表（如果存在）。
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	// 创建 User 表。
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// 查询 User 表中的行数。
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var count int // 用于存储查询结果
	// 检查查询是否成功，以及查询结果是否为 0。
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err) // 如果检查失败，终止测试并报错
	}
}
