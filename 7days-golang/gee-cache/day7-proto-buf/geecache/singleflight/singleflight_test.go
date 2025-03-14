package singleflight

import (
	"testing"
)

// TestDo 是一个单元测试，用于验证 singleflight.Group 的 Do 方法是否正确工作。
func TestDo(t *testing.T) {
	// 创建一个 singleflight.Group 实例
	var g Group

	// 调用 Do 方法，执行一个简单的计算函数
	// 计算函数返回字符串 "bar"，并返回 nil 作为错误
	v, err := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})

	// 验证 Do 方法的返回值是否符合预期
	if v != "bar" || err != nil {
		// 如果返回值或错误不符合预期，记录错误信息
		t.Errorf("Do v = %v, error = %v", v, err)
	}
}
