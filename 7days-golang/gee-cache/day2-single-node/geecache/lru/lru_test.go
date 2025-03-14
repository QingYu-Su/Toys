package lru

import (
	"reflect" // 用于比较切片内容
	"testing" // Go 的单元测试框架
)

// 定义一个简单的 String 类型，用于测试缓存值
type String string

// 实现 Value 接口的 Len() 方法，计算字符串的长度
func (d String) Len() int {
	return len(d)
}

// 测试 Get 方法是否正确
func TestGet(t *testing.T) {
	lru := New(int64(0), nil)       // 创建一个无大小限制的 LRU 缓存
	lru.Add("key1", String("1234")) // 添加一个键值对

	// 测试缓存命中
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed") // 如果未命中或值不正确，测试失败
	}

	// 测试缓存未命中
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed") // 如果错误地命中了不存在的键，测试失败
	}
}

// 测试 RemoveOldest 方法是否正确移除最久未使用的缓存项
func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"     // 定义三个键
	v1, v2, v3 := "value1", "value2", "v3" // 定义对应的值
	cap := len(k1 + k2 + v1 + v2)          // 计算缓存容量，刚好容纳两个键值对

	lru := New(int64(cap), nil) // 创建一个有大小限制的 LRU 缓存
	lru.Add(k1, String(v1))     // 添加第一个键值对
	lru.Add(k2, String(v2))     // 添加第二个键值对
	lru.Add(k3, String(v3))     // 添加第三个键值对，触发移除最久未使用的键值对

	// 验证 key1 是否被正确移除
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed") // 如果 key1 未被移除或缓存大小不正确，测试失败
	}
}

// 测试 OnEvicted 回调函数是否正确触发
func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0) // 用于存储被移除的键
	// 定义回调函数，将被移除的键添加到 keys 切片中
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback)   // 创建一个有大小限制的 LRU 缓存，并设置回调函数
	lru.Add("key1", String("123456")) // 添加第一个键值对，超出缓存限制
	lru.Add("k2", String("k2"))       // 添加第二个键值对，触发移除第一个键值对
	lru.Add("k3", String("k3"))       // 添加第三个键值对，触发移除第二个键值对
	lru.Add("k4", String("k4"))       // 添加第四个键值对，触发移除第三个键值对

	// 验证被移除的键是否正确
	expect := []string{"key1", "k2"} // 预期被移除的键
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect) // 如果实际结果与预期不符，测试失败
	}
}

// 测试 Add 方法是否正确更新缓存大小
func TestAdd(t *testing.T) {
	lru := New(int64(0), nil)     // 创建一个无大小限制的 LRU 缓存
	lru.Add("key", String("1"))   // 添加第一个键值对
	lru.Add("key", String("111")) // 更新同一个键的值

	// 验证缓存大小是否正确更新
	if lru.nbytes != int64(len("key")+len("111")) {
		t.Fatal("expected 6 but got", lru.nbytes) // 如果缓存大小不正确，测试失败
	}
}
