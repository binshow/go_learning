package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_learning/7days-golang/GeeRPC/day01-codec/codec"
	"go_learning/7days-golang/GeeRPC/day04-timeout/server"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// -------------------------------------------
// @file          : client.go
// @author        : binshow
// @time          : 2022/6/9 9:58 AM
// @description   :
// -------------------------------------------



// Call 来承载一次RPC所需要的信息
type Call struct {
	Seq  			uint64
	ServiceMethod  	string
	Args			interface{}
	Reply 			interface{}
	Error           error
	Done     		chan *Call // 为了支持异步调用，当调用结束时，会调用 call.done 通知调用方
}

func (call *Call) done() {
	call.Done <- call
}


type Client struct {
	cc    		codec.Codec			// 编码器，序列化发送出去的请求以及序列化接收到的response
	opt   	 	*server.Option
	sending  	sync.Mutex			// 互斥锁，保证消息的有序发送
	header   	codec.Header		// request 的请求头信息
	mu 			sync.Mutex
	seq         uint64
	pending     map[uint64]*Call // 存储未处理完的请求
	closing     bool             // user has called Close
	shutdown    bool             // server has told us to stop ， 一般是有错误发生
}
var _ io.Closer = (*Client)(nil)
var ErrShutdown = errors.New("connection is shut down")

func (client *Client) Close() error{
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}


func (client *Client) IsAvailable() bool{
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

//服务端或客户端发生错误时调用，将 shutdown 设置为 true，且将错误信息通知所有 pending 状态的 call
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}


func parseOptions(opts ...*server.Option) (*server.Option, error) {
	// if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil {
		return server.DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = server.DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = server.DefaultOption.CodecType
	}
	return opt, nil
}




// Dial connects to an RPC server at the specified network address
func Dial(network, address string, opts ...*server.Option) (client *Client, err error) {
	return dialTimeout(NewClient , network , address , opts...)
}

type clientResult struct {
	client *Client
	err    error
}

type newClientFunc func(conn net.Conn , opt *server.Option) (client *Client , err error)

func dialTimeout(f newClientFunc , network , address string , opts ...*server.Option) (client *Client , err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	// 1. 处理连接超时
	conn, err := net.DialTimeout(network , address , opt.ConnectTimeout)
	if err != nil {
		return nil, err
	}
	// close the connection if client is nil
	defer func() {
		if client == nil {
			_ = conn.Close()
		}
	}()

	ch := make(chan clientResult)
	// 2. 使用 子 goroutine 执行 NewClient。执行完之后通过 channel 发送结果
	go func() {
		client , err := f(conn , opt)
		ch <- clientResult{
			client: client,
			err:    err,
		}
	}()

	if opt.ConnectTimeout == 0 {
		result := <- ch
		return result.client , result.err
	}

	// 如果到了指定的超时时间，channel中还没有结果，就会报错
	select {
	case <- time.After(opt.ConnectTimeout):
		return nil , fmt.Errorf("rpc client connect timeout: expect within %s" , opt.ConnectTimeout)
	case result := <-ch:
		return result.client , result.err
	}
}



// NewClient 创建Client 实例之后，首先要完成一开始的协议交换，即发送 Option 信息给服务端
func NewClient(conn net.Conn, opt *server.Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	// send options with server
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *server.Option) *Client {
	client := &Client{
		seq:     1, // seq starts with 1, 0 means invalid call
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	// 再创建一个子goroutine 接收服务端的resp
	go client.receive()
	return client
}




func (client *Client) send(call *Call) {
	// make sure that the client will send a complete request
	client.sending.Lock()
	defer client.sending.Unlock()

	// register this call.
	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	// prepare request header
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	// encode and send the request
	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq)
		// call may be nil, it usually means that Write partially failed,
		// client has received the response and handled
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// 接收resp
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		call := client.removeCall(h.Seq)
		switch {
		case call == nil:
			// it usually means that Write partially failed
			// and call was already removed.
			err = client.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	// error occurs, so terminateCalls pending calls
	client.terminateCalls(err)
}

// Go 暴露给用户的RPC调用接口，是一个异步的接口，返回Call实例
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

//// Call 暴露给用户的RPC调用接口，是一个同步的接口
//func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
//	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
//	return call.Error
//}

// Call 暴露给用户的RPC调用接口，是一个同步的接口
// 使用context包实现，控制权交给用户，控制更为灵活
// 用户可以使用 context.WithTimeout 创建具备超时检测能力的 context 对象来控制
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	call :=  client.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}