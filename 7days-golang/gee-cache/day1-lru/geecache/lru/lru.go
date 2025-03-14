package lru

import "container/list" // 导入标准库中的双向链表包，用于实现LRU缓存的核心数据结构

// Cache 是一个 LRU 缓存结构，包含以下字段：
type Cache struct {
	maxBytes  int64                         // 缓存的最大字节数限制
	nbytes    int64                         // 当前缓存中存储的数据总字节数
	ll        *list.List                    // 双向链表，用于记录缓存项的访问顺序
	cache     map[string]*list.Element      // 哈希表，用于快速查找缓存项
	OnEvicted func(key string, value Value) // 回调函数，当缓存项被移除时调用
}

// entry 是缓存项的结构，包含键和值
type entry struct {
	key   string // 缓存项的键
	value Value  // 缓存项的值，Value 是一个接口类型
}

// Value 是一个接口，要求缓存值必须实现 Len() 方法，用于计算值的大小（字节数）
type Value interface {
	Len() int
}

// New 函数用于创建一个新的 LRU 缓存实例
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,                       // 设置缓存的最大字节数限制
		ll:        list.New(),                     // 初始化双向链表
		cache:     make(map[string]*list.Element), // 初始化哈希表
		OnEvicted: onEvicted,                      // 设置缓存项被移除时的回调函数
	}
}

// Get 方法用于从缓存中获取一个键的值
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok { // 如果键存在于缓存中
		c.ll.MoveToFront(ele)    // 将对应的链表元素移动到链表头部（表示最近访问）
		kv := ele.Value.(*entry) // 获取缓存项的值
		return kv.value, true    // 返回值和 true 表示查找成功
	}
	return nil, false // 如果键不存在，返回 nil 和 false
}

// RemoveOldest 方法用于移除最久未使用的缓存项
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() // 获取链表尾部的元素（最久未使用的缓存项）
	if ele != nil {
		c.ll.Remove(ele)                                       // 从链表中移除该元素
		kv := ele.Value.(*entry)                               // 获取缓存项的键和值
		delete(c.cache, kv.key)                                // 从哈希表中删除对应的键
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) // 更新当前缓存的总字节数
		if c.OnEvicted != nil {                                // 如果设置了回调函数
			c.OnEvicted(kv.key, kv.value) // 调用回调函数，通知缓存项被移除
		}
	}
}

// Add 方法用于向缓存中添加一个键值对
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // 如果键已经存在
		c.ll.MoveToFront(ele)                                  // 将对应的链表元素移动到链表头部
		kv := ele.Value.(*entry)                               // 获取缓存项
		c.nbytes += int64(value.Len()) - int64(kv.value.Len()) // 更新当前缓存的总字节数
		kv.value = value                                       // 更新缓存项的值
	} else {
		ele := c.ll.PushFront(&entry{key, value})        // 在链表头部插入新的缓存项
		c.cache[key] = ele                               // 在哈希表中添加键和对应的链表元素
		c.nbytes += int64(len(key)) + int64(value.Len()) // 更新当前缓存的总字节数
	}

	// 如果缓存的总字节数超过了限制，则移除最久未使用的缓存项
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len 方法返回当前缓存中存储的键值对数量
func (c *Cache) Len() int {
	return c.ll.Len() // 返回双向链表的长度
}
