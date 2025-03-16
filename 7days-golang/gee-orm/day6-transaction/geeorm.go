package geeorm

import (
	"database/sql" // 提供数据库操作功能

	"geeorm/dialect"
	"geeorm/log"     // 自定义日志包，用于记录日志
	"geeorm/session" // 自定义 session 包，用于封装数据库操作
)

// Engine 是 GeeORM 的核心结构体，封装了数据库连接。
type Engine struct {
	db      *sql.DB // 数据库连接对象
	dialect dialect.Dialect
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

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	// 创建 Engine 实例并返回。
	e = &Engine{db: db, dialect: dial}
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
	return session.New(engine.db, engine.dialect)
}

// TxFunc 是一个事务函数类型，用于在事务中执行操作。
// 它接受一个 *session.Session 参数，并返回一个结果和一个错误。
type TxFunc func(*session.Session) (interface{}, error)

// Transaction 执行一个事务函数，并确保事务的正确提交或回滚。
// 参数：
//   - f：事务函数，需要在事务中执行的操作。
//
// 返回值：
//   - result：事务函数的返回结果。
//   - err：事务执行过程中发生的错误。
func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession() // 创建一个新的 Session 实例
	if err := s.Begin(); err != nil {
		return nil, err // 如果事务开始失败，直接返回错误
	}

	defer func() {
		// 捕获 panic，确保事务回滚
		if p := recover(); p != nil {
			_ = s.Rollback() // 回滚事务
			panic(p)         // 重新抛出 panic
		} else if err != nil {
			_ = s.Rollback() // 如果有错误，回滚事务
		} else {
			err = s.Commit() // 如果没有错误，提交事务
		}
	}()

	// 执行事务函数
	return f(s)
}
