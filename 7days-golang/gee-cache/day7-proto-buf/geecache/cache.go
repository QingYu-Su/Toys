package geecache

import (
	"geecache/lru" // 引入本地的 LRU 缓存实现包
	"sync"         // 引入 sync 包，用于线程安全的互斥锁
)

// cache 是一个封装了 LRU 缓存的结构体，用于实现线程安全的缓存操作。
type cache struct {
	mu         sync.Mutex // 互斥锁，用于保护对缓存的并发访问
	lru        *lru.Cache // LRU 缓存实例
	cacheBytes int64      // 缓存的最大字节数限制
}

// add 方法用于向缓存中添加一个键值对。
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()         // 加锁，确保线程安全
	defer c.mu.Unlock() // 确保在函数退出时释放锁

	// 如果 LRU 缓存实例尚未初始化，则创建一个新的缓存实例
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	// 将键值对添加到 LRU 缓存中
	c.lru.Add(key, value)
}

// get 方法用于从缓存中获取一个键对应的值。
func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()         // 加锁，确保线程安全
	defer c.mu.Unlock() // 确保在函数退出时释放锁

	// 如果 LRU 缓存实例尚未初始化，直接返回
	if c.lru == nil {
		return
	}

	// 从 LRU 缓存中获取键对应的值
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok // 将值断言为 ByteView 类型并返回
	}

	return // 如果键不存在，返回默认值
}
