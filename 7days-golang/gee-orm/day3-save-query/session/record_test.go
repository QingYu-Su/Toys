package session

import "testing"

// 测试用的 User 结构体
var (
	user1 = &User{"Tom", 18}  // 第一个测试用户
	user2 = &User{"Sam", 25}  // 第二个测试用户
	user3 = &User{"Jack", 25} // 第三个测试用户
)

// testRecordInit 初始化测试环境，创建表并插入测试记录。
func testRecordInit(t *testing.T) *Session {
	t.Helper()                       // 标记为辅助函数，便于调试和日志记录
	s := NewSession().Model(&User{}) // 创建一个新的 Session 实例，并设置模型为 User
	// 删除表（如果存在），创建表，并插入两条测试记录
	err1 := s.DropTable()             // 删除表（如果存在）
	err2 := s.CreateTable()           // 创建表
	_, err3 := s.Insert(user1, user2) // 插入两条测试记录
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records") // 如果初始化失败，终止测试并报错
	}
	return s // 返回初始化后的 Session 实例
}

// TestSession_Insert 测试 Session 的 Insert 方法。
func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)           // 初始化测试环境
	affected, err := s.Insert(user3) // 插入第三条测试记录
	if err != nil || affected != 1 { // 检查是否成功插入一条记录
		t.Fatal("failed to create record") // 如果插入失败，终止测试并报错
	}
}

// TestSession_Find 测试 Session 的 Find 方法。
func TestSession_Find(t *testing.T) {
	s := testRecordInit(t) // 初始化测试环境
	var users []User       // 用于存储查询结果的切片
	// 查询所有记录并填充到 users 切片中
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all") // 如果查询失败或结果数量不正确，终止测试并报错
	}
}
