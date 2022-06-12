package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go_learning/thrift_study/idl/gen-go/Sample"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

//定义服务
type GreeterService struct {
}

//实现IDL里定义的接口 Sample.GreeterService
var _ Sample.Greeter = (*GreeterService)(nil)

//SayHello
func (this *GreeterService) SayHello(ctx context.Context, u *Sample.User) (r *Sample.Response, err error) {
	strJson, _ := json.Marshal(u)
	return &Sample.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

//GetUser
func (this *GreeterService) GetUser(ctx context.Context, uid int32) (r *Sample.Response, err error) {
	return &Sample.Response{ErrCode: 1, ErrMsg: "user not exist."}, nil
}



/**
	thrift 中分层如下：
	1. 传输层 transport layer : 负责从网络中读取和写入数据，定义了具体的网络传输协议比如tcp 等
	2. 协议层 protocol layer  : 定义了数据传输的格式，负责网络传输数据的序列化方式，比如 json、xml、二进制等
	3. 处理层 processor layer : 指的就是 idl 文件生成的，封装了底层网络传输和序列化方式，并委托给用户生成的 handler 来具体实现
	4. 服务层 service layer   : 整合上述的组件，提供具体的网络线程 IO模型等


 */


func main() {
	//命令行参数
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	flag.Parse()

	//protocol：协议层
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	//buffered
	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	//framed
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	//handler
	handler := &GreeterService{}

	//transport,no secure 传输层
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Println("error running server:", err)
	}

	//processor 处理层
	processor := Sample.NewGreeterProcessor(handler)
	fmt.Println("Starting the simple server... on ", *addr)


	// start tcp server
	// TSimpleServer 简单的单线程服务模型
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}
}
