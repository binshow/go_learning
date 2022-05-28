package frame

import (
	"encoding/binary"
	"errors"
	"io"
)


const (
	CommandConn   = iota + 0x01 // 0x01，连接请求包
	CommandSubmit               // 0x02，消息请求包
)

const (
	CommandConnAck   = iota + 0x81 // 0x81，连接请求的响应包
	CommandSubmitAck               // 0x82，消息请求的响应包
)

type FramePayload []byte

// StreamFrameCodec 统一的接口类型
type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error // data -> frame，并写入io.Writer
	Decode(io.Reader)(FramePayload , error)  // 从io.Reader中提取frame payload，并返回给上层
}

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct {

}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, framePayload FramePayload) error {
	var f = framePayload
	var totalLen int32 = int32(len(framePayload) + 4)

	err := binary.Write(w , binary.BigEndian , &totalLen)
	if err != nil {
		return err
	}

	// write the frame payload to outbound stream
	n , err := w.Write([]byte(f))
	if err != nil {
		return err
	}

	if n != len(framePayload) {
		return ErrShortWrite
	}
	return nil
}

func (p *myFrameCodec) Decode(r io.Reader)(FramePayload , error) {
	var totalLen int32
	err := binary.Read(r , binary.BigEndian , &totalLen)
	if err != nil {
		return nil, err
	}
	buf := make([]byte , totalLen - 4)
	n , err := io.ReadFull(r , buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil , ErrShortRead
	}

	return FramePayload(buf) , nil
}