package main

import (
	"context"
	client "go_learning/7days-golang/GeeRPC/day05-http-debug/client"
	geerpc "go_learning/7days-golang/GeeRPC/day05-http-debug/server"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/8 11:51 PM
// @description   : 超时处理时RPC框架的一个比较基本的能力，如果缺少超时处理机制，无论是服务端还是客户端都容易因为网络或其他错误导致挂死，资源耗尽
// -------------------------------------------

/*
纵观整个远程调用的过程，需要客户端处理超时的地方有：
1. 与服务端建立连接，导致的超时
2. 发送请求到服务端，写报文导致的超时
3. 等待服务端处理时，等待处理导致的超时（比如服务端已挂死，迟迟不响应）
4. 从服务端接收响应时，读报文导致的超时

需要服务端处理超时的地方有：
1. 读取客户端请求报文时，读报文导致的超时
2. 发送响应报文时，写报文导致的超时
3. 调用映射服务的方法时，处理报文导致的超时

GeeRPC 在 3 个地方添加了超时处理机制。分别是：
1）客户端创建连接时
2）客户端 Client.Call() 整个过程导致的超时（包含发送报文，等待处理，接收报文所有阶段）
3）服务端处理报文，即 Server.handleRequest 超时。

*/

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addrCh chan string) {
	var foo Foo
	l, _ := net.Listen("tcp", ":9999")
	_ = geerpc.Register(&foo)
	geerpc.HandleHTTP()
	addrCh <- l.Addr().String()
	_ = http.Serve(l, nil)
}

func call(addrCh chan string) {
	client, _ := client.DialHTTP("tcp", <-addrCh)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}

func main() {
	log.SetFlags(0)
	ch := make(chan string)
	go call(ch)
	startServer(ch)
}