package session

import (
	"database/sql" // 提供数据库操作功能
	"geeorm/log"   // 引入自定义的日志包，用于记录日志
	"strings"      // 提供字符串操作功能
)

// Session 是一个封装了数据库操作的结构体，用于简化 SQL 操作。
type Session struct {
	db      *sql.DB         // 数据库连接对象
	sql     strings.Builder // 用于构建 SQL 语句
	sqlVars []interface{}   // 用于存储 SQL 语句的参数
}

// New 创建一个新的 Session 实例。
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear 清空当前构建的 SQL 语句和参数。
func (s *Session) Clear() {
	s.sql.Reset()   // 清空 SQL 语句
	s.sqlVars = nil // 清空 SQL 参数
}

// DB 返回当前 Session 的数据库连接对象。
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw 用于构建原始的 SQL 语句和参数。
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)                   // 将 SQL 语句追加到 Builder 中
	s.sql.WriteString(" ")                   // 在 SQL 语句后追加一个空格（可选）
	s.sqlVars = append(s.sqlVars, values...) // 将参数追加到 sqlVars 中
	return s
}

// Exec 执行当前构建的 SQL 语句，并返回执行结果。
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()                     // 执行完成后清空 SQL 语句和参数
	log.Info(s.sql.String(), s.sqlVars) // 记录执行的 SQL 语句和参数
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err) // 如果执行失败，记录错误日志
	}
	return
}

// QueryRow 查询单行数据。
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()                     // 查询完成后清空 SQL 语句和参数
	log.Info(s.sql.String(), s.sqlVars) // 记录执行的 SQL 语句和参数
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows 查询多行数据。
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()                     // 查询完成后清空 SQL 语句和参数
	log.Info(s.sql.String(), s.sqlVars) // 记录执行的 SQL 语句和参数
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err) // 如果查询失败，记录错误日志
	}
	return
}
