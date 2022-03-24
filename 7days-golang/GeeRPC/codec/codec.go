package codec

import "io"

/**
	一个典型的RPC调用如下所示：
	err = client.Call("Arith.Multiply", args, &reply)
	其中客户端发送的请求中应该包括 服务名、方法名、参数这三个
	服务端的响应中应该包含	错误、返回值 这两个内容

	那么我们将请求和响应中的参数和返回值抽象为 body，
	剩余的信息放在 header 中，那么就可以抽象出数据结构 Header
 */


type Header struct {
	ServiceMethod string // format "Service.Method" 服务名和方法名
	Seq           uint64 // sequence number chosen by client，请求的 ID，用来区分不同的请求
	Error         string
}


//Codec 是对消息体进行编码和解码的接口，抽象出了结构就可以实现 不同的编码和解码实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error         // 解析header
	ReadBody(interface{}) error       // 解析body
	Write(*Header, interface{}) error // 写数据
}


// NewCodecFunc 抽象出 Codec 的构造函数，客户端和服务端可以通过 Codec 的 Type 得到构造函数
type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)


//NewCodecFuncMap 构造一个map，key为类型，value为对应解码器的构造函数，类似于工厂模式
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
