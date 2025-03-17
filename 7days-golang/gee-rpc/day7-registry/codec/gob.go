package codec

import (
	"bufio"        // 用于创建带缓冲的 I/O 流
	"encoding/gob" // 提供 Gob 编解码功能
	"io"           // 提供基本的 I/O 操作接口
	"log"          // 用于记录日志信息
)

// GobCodec 是一个基于 Gob 编解码的 Codec 实现
type GobCodec struct {
	conn io.ReadWriteCloser // 底层的读写关闭接口，用于网络通信
	buf  *bufio.Writer      // 带缓冲的写入器，用于提高写入效率
	dec  *gob.Decoder       // Gob 解码器，用于从 conn 中读取数据并解码
	enc  *gob.Encoder       // Gob 编码器，用于将数据编码后写入 buf
}

// 确保 GobCodec 实现了 Codec 接口
var _ Codec = (*GobCodec)(nil)

// NewGobCodec 创建一个新的 GobCodec 实例
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn) // 创建带缓冲的写入器
	return &GobCodec{
		conn: conn,                 // 设置底层的读写关闭接口
		buf:  buf,                  // 设置带缓冲的写入器
		dec:  gob.NewDecoder(conn), // 创建 Gob 解码器
		enc:  gob.NewEncoder(buf),  // 创建 Gob 编码器
	}
}

// ReadHeader 从底层连接中读取头部信息并解码到 Header 结构体中
func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h) // 使用 Gob 解码器解码头部信息
}

// ReadBody 从底层连接中读取正文内容并解码到指定的 body 接口类型中
func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body) // 使用 Gob 解码器解码正文内容
}

// Write 将头部信息和正文内容编码后写入底层连接
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush() // 确保缓冲区中的数据被写入底层连接
		if err != nil {
			_ = c.Close() // 如果发生错误，关闭连接
		}
	}()
	if err := c.enc.Encode(h); err != nil { // 编码头部信息
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil { // 编码正文内容
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

// Close 关闭底层的读写关闭接口
func (c *GobCodec) Close() error {
	return c.conn.Close() // 关闭底层连接
}
