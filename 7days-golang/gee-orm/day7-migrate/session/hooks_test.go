package session

import (
	"geeorm/log"
	"testing"
)

// Account 是一个测试用的结构体，模拟数据库中的用户账户表。
type Account struct {
	ID       int    `geeorm:"PRIMARY KEY"` // 主键字段
	Password string // 密码字段
}

// BeforeInsert 是 Account 的钩子方法，在插入记录前被调用。
func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", account) // 记录日志
	account.ID += 1000                 // 修改 ID，模拟自定义逻辑
	return nil
}

// AfterQuery 是 Account 的钩子方法，在查询记录后被调用。
func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account) // 记录日志
	account.Password = "******"      // 修改密码字段，隐藏真实密码
	return nil
}

// TestSession_CallMethod 测试 Session 的 CallMethod 方法是否能够正确调用钩子方法。
func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})                           // 创建 Session 并设置模型为 Account
	_ = s.DropTable()                                             // 删除表（如果存在）
	_ = s.CreateTable()                                           // 创建表
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"}) // 插入两条测试记录

	u := &Account{} // 用于存储查询结果的结构体实例

	// 查询第一条记录
	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u) // 如果查询失败或钩子方法未正确执行，终止测试
	}
}
