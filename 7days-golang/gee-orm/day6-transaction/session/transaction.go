package session

import "geeorm/log"

// Begin 开始一个新的数据库事务。
// 返回值：
//   - err：如果开始事务时发生错误，返回错误信息。
func (s *Session) Begin() (err error) {
	log.Info("transaction begin") // 记录事务开始的日志
	// 调用数据库连接的 Begin 方法启动事务
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err) // 如果启动事务失败，记录错误日志
		return
	}
	return
}

// Commit 提交当前事务。
// 返回值：
//   - err：如果提交事务时发生错误，返回错误信息。
func (s *Session) Commit() (err error) {
	log.Info("transaction commit") // 记录事务提交的日志
	// 调用事务的 Commit 方法提交事务
	if err = s.tx.Commit(); err != nil {
		log.Error(err) // 如果提交失败，记录错误日志
	}
	return
}

// Rollback 回滚当前事务。
// 返回值：
//   - err：如果回滚事务时发生错误，返回错误信息。
func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback") // 记录事务回滚的日志
	// 调用事务的 Rollback 方法回滚事务
	if err = s.tx.Rollback(); err != nil {
		log.Error(err) // 如果回滚失败，记录错误日志
	}
	return
}
