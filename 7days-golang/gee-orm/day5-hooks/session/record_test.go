package session

import "testing"

// 测试用的 User 结构体实例
var (
	user1 = &User{"Tom", 18} // 第一个测试用户
	user2 = &User{"Sam", 25} // 第二个测试用户
)

// testRecordInit 初始化测试环境，创建表并插入测试记录。
func testRecordInit(t *testing.T) *Session {
	t.Helper()                       // 标记为辅助函数
	s := NewSession().Model(&User{}) // 创建 Session 并设置模型为 User
	// 初始化测试环境：删除表（如果存在）、创建表、插入测试记录
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records") // 如果初始化失败，终止测试
	}
	return s
}

// TestSession_Limit 测试 Limit 方法是否能正确限制查询结果的数量。
func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t) // 初始化测试环境
	var users []User       // 用于存储查询结果
	// 使用 Limit 方法限制查询结果为 1 条
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition") // 如果查询失败或结果数量不正确，终止测试
	}
}

// TestSession_Update 测试 Update 方法是否能正确更新记录。
func TestSession_Update(t *testing.T) {
	s := testRecordInit(t) // 初始化测试环境
	// 使用 Where 方法设置条件，然后调用 Update 方法更新记录
	affected, _ := s.Where("Name = ?", "Tom").Update("Age", 30)
	// 查询更新后的记录
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	// 检查是否成功更新了 1 条记录，且更新后的年龄是否为 30
	if affected != 1 || u.Age != 30 {
		t.Fatal("failed to update") // 如果更新失败，终止测试
	}
}

// TestSession_DeleteAndCount 测试 Delete 和 Count 方法是否能正确删除记录并统计数量。
func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t) // 初始化测试环境
	// 使用 Where 方法设置条件，然后调用 Delete 方法删除记录
	affected, _ := s.Where("Name = ?", "Tom").Delete()
	// 查询剩余记录的数量
	count, _ := s.Count()

	// 检查是否成功删除了 1 条记录，且剩余记录数量是否为 1
	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count") // 如果删除或统计失败，终止测试
	}
}
