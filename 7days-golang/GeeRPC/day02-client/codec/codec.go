package codec

import "io"

// -------------------------------------------
// @file          : codec.go
// @author        : binshow
// @time          : 2022/6/8 11:52 PM
// @description   : 一个典型的 RPC 调用如下： err = client.Call("Arith.Multiply", args, &reply)
//					需要包括： 服务名，方法名，参数三个，  服务端的响应 包含 reply 和 err
// -------------------------------------------

// 将request 和 response 中的参数 和 返回值 抽象成 body，剩余的信息就可以放在header中了

type Header struct {
	ServiceMethod 	string		//服务名和方法名，通常与 Go 语言中的结构体和方法相映射。
	Seq				uint64		//请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求。
	Error           string		//是错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
}

// Codec 抽象出对消息body就那些编解码的接口，这样可以实现不同的Codec实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header , interface{}) error
}

// NewCodecFunc 抽象出Codec的构造函数
type NewCodecFunc  func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap  map[Type]NewCodecFunc		// 通过一个map来保存 不同的Type对应的构造函数，类似于工厂模式

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}






