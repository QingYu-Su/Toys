package geeorm

import (
	"database/sql" // 提供数据库操作功能

	"geeorm/log"     // 自定义日志包，用于记录日志
	"geeorm/session" // 自定义 session 包，用于封装数据库操作
)

// Engine 是 GeeORM 的核心结构体，封装了数据库连接。
type Engine struct {
	db *sql.DB // 数据库连接对象
}

// NewEngine 创建一个新的 Engine 实例，并连接到指定的数据库。
func NewEngine(driver, source string) (e *Engine, err error) {
	// 使用 sql.Open 创建数据库连接。
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err) // 如果连接失败，记录错误日志
		return
	}

	// 使用 db.Ping 检查数据库连接是否成功。
	if err = db.Ping(); err != nil {
		log.Error(err) // 如果连接失败，记录错误日志
		return
	}

	// 创建 Engine 实例并返回。
	e = &Engine{db: db}
	log.Info("Connect database success") // 记录成功连接数据库的日志
	return
}

// Close 关闭 Engine 中的数据库连接。
func (engine *Engine) Close() {
	// 尝试关闭数据库连接。
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database") // 如果关闭失败，记录错误日志
	}
	log.Info("Close database success") // 记录成功关闭数据库的日志
}

// NewSession 创建一个新的 Session 实例，用于执行数据库操作。
func (engine *Engine) NewSession() *session.Session {
	// 使用 session.New 创建一个新的 Session 实例，并传递当前 Engine 的数据库连接。
	return session.New(engine.db)
}
