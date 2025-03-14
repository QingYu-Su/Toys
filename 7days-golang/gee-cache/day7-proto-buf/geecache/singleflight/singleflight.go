package singleflight

import "sync"

// call 是一个结构体，用于存储计算结果和同步状态。
type call struct {
	wg  sync.WaitGroup // 用于同步等待计算完成
	val interface{}    // 存储计算结果
	err error          // 存储计算过程中可能发生的错误
}

// Group 是一个单例模式的结构体，用于管理并发请求。
type Group struct {
	mu sync.Mutex       // 互斥锁，保护 map 的并发访问
	m  map[string]*call // 存储每个 key 对应的 call 结构体
}

// Do 方法用于执行计算逻辑，并确保对于同一个 key，只执行一次计算。
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock() // 加锁，确保对 map 的访问是线程安全的
	if g.m == nil {
		g.m = make(map[string]*call) // 如果 map 未初始化，则初始化
	}
	if c, ok := g.m[key]; ok { // 检查是否已经存在一个 call 实例
		g.mu.Unlock()       // 如果存在，释放锁
		c.wg.Wait()         // 等待计算完成
		return c.val, c.err // 返回计算结果
	}

	// 如果不存在，创建一个新的 call 实例
	c := new(call)
	c.wg.Add(1)   // 增加 WaitGroup 的计数
	g.m[key] = c  // 将 call 实例存储到 map 中
	g.mu.Unlock() // 释放锁

	// 执行计算逻辑
	c.val, c.err = fn()
	c.wg.Done() // 通知等待的 goroutine 计算已完成

	g.mu.Lock()      // 再次加锁
	delete(g.m, key) // 删除 map 中的 call 实例
	g.mu.Unlock()    // 释放锁

	return c.val, c.err // 返回计算结果
}
