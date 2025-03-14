package consistenthash

import (
	"hash/crc32" // 导入 CRC32 哈希算法
	"sort"       // 导入 sort 包，用于对哈希值进行排序
	"strconv"    // 导入 strconv 包，用于整数和字符串的转换
)

// Hash 是哈希函数的类型，定义为一个函数，接受字节切片并返回一个 32 位无符号整数。
type Hash func(data []byte) uint32

// Map 是一致性哈希的实现，包含哈希函数、虚拟节点数量、哈希值列表和哈希值到键的映射。
type Map struct {
	hash     Hash           // 哈希函数，用于计算哈希值
	replicas int            // 虚拟节点的数量，用于提高负载均衡性
	keys     []int          // 存储所有哈希值的有序列表
	hashMap  map[int]string // 哈希值到键的映射，用于快速查找
}

// New 创建一个新的 Map 实例，接受虚拟节点数量和哈希函数作为参数。
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,             // 设置虚拟节点数量
		hash:     fn,                   // 设置哈希函数
		hashMap:  make(map[int]string), // 初始化哈希映射
	}

	// 如果未提供哈希函数，则默认使用 CRC32 的 IEEE 标准实现
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// Add 将一组键（通常是节点或服务器）添加到一致性哈希环中。
func (m *Map) Add(keys ...string) {
	for _, key := range keys { // 遍历所有传入的键
		for i := 0; i < m.replicas; i++ { // 为每个键生成多个虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + key))) // 计算虚拟节点的哈希值
			m.keys = append(m.keys, hash)                      // 将哈希值添加到列表中
			m.hashMap[hash] = key                              // 将哈希值映射到键
		}
	}

	sort.Ints(m.keys) // 对哈希值列表进行排序，以便后续查找
}

// Get 根据给定的键查找对应的节点。
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 { // 如果哈希环为空，直接返回空字符串
		return ""
	}

	hash := int(m.hash([]byte(key))) // 计算键的哈希值

	// 使用 sort.Search 查找第一个大于或等于该哈希值的索引
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// 如果索引超出范围，则通过取模操作实现环形查找
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
