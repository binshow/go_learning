package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

// -------------------------------------------
// @file          : gob.go
// @author        : binshow
// @time          : 2022/6/9 12:01 AM
// @description   : GobCodec 实现了 Codec 编解码的接口
// -------------------------------------------


type GobCodec struct {
	conn io.ReadWriteCloser		// conn 是由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例
	buf  *bufio.Writer		 	// buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能。
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// 下面这行代码是确保 接口被实现常用的方式，使用强制类型转换 确保 GobCodec 实现了 Codec 的接口，这样编译期就可以发现
// 将空值 nil 转换成 *GobCodec 类型，再转换成 Codec 接口类型变量！！！
var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}