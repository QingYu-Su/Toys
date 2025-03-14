package consistenthash

import (
	"strconv" // 用于字符串和整数的转换
	"testing" // 用于编写和运行测试
)

// TestHashing 测试一致性哈希算法的实现是否正确。
func TestHashing(t *testing.T) {
	// 创建一个新的一致性哈希实例，使用自定义的哈希函数。
	// 自定义哈希函数将键转换为整数，并直接返回该整数作为哈希值。
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key)) // 将键转换为整数
		return uint32(i)                  // 返回整数作为哈希值
	})

	// 添加节点到一致性哈希环中。
	// 给定的哈希函数会为每个节点生成多个虚拟节点。
	// 例如，节点 "6" 会生成虚拟节点 0+6=6, 1+6=16, 2+6=26。
	// 因此，添加 "6", "4", "2" 后，哈希环中的哈希值为：
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	// 定义测试用例，测试键是否被正确映射到节点。
	testCases := map[string]string{
		"2":  "2", // 键 "2" 应该映射到节点 "2"
		"11": "2", // 键 "11" 应该映射到节点 "2"（因为 11 在 2 和 12 之间）
		"23": "4", // 键 "23" 应该映射到节点 "4"（因为 23 在 22 和 24 之间）
		"27": "2", // 键 "27" 应该映射到节点 "2"（因为 27 在 26 和 2 之间，环形查找）
	}

	// 遍历测试用例，验证每个键是否被正确映射。
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// 添加新的节点 "8"，生成虚拟节点 8, 18, 28。
	hash.Add("8")

	// 更新测试用例，键 "27" 现在应该映射到节点 "8"。
	testCases["27"] = "8"

	// 再次验证每个键是否被正确映射。
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
