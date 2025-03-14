package geecache

// ByteView 是一个结构体，用于表示一个不可变的字节切片视图。
// 它封装了字节切片，防止外部直接修改数据。
type ByteView struct {
	b []byte // 内部存储的字节切片
}

// Len 方法返回字节切片的长度。
func (v ByteView) Len() int {
	return len(v.b) // 返回字节切片的长度
}

// ByteSlice 方法返回一个字节切片的副本，而不是原始切片。
// 这样可以防止外部代码直接修改内部数据。
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b) // 调用 cloneBytes 函数复制字节切片
}

// String 方法将字节切片转换为字符串。
func (v ByteView) String() string {
	return string(v.b) // 将字节切片转换为字符串
}

// cloneBytes 是一个辅助函数，用于复制一个字节切片。
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b)) // 创建一个新的字节切片，长度与原始切片相同
	copy(c, b)                // 将原始切片的内容复制到新切片中
	return c                  // 返回复制后的切片
}
