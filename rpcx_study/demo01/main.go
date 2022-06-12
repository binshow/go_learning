package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/12 4:56 PM
// @description   : quickstart
// -------------------------------------------


var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	go startServer()

	time.Sleep(time.Second * 1)

	//1. 使用最简单的 点对点 ，客户端直连服务器来获取服务器地址
	d , _:= client.NewPeer2PeerDiscovery("tcp@" + *addr , "")

	//2. 创建客户端 ，client.Failtry 表示如何处理调用失败
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	//3. 定义了请求
	args := &Args{
		A: 10,
		B: 20,
	}

	//4. 定义了 响应对象
	reply := &Reply{}

	//5. 调用远程服务并且 同步获取了结果
	//err := xClient.Call(context.Background(), "Mul", args, reply)
	//if err != nil {
	//	log.Fatal("failed to call : %v" , err)
	//}
	//fmt.Printf("%d * %d = %d", args.A, args.B, reply.C)


	//6. 异步调用
	call, err := xClient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatal("failed to call : %v" , err)
	}

	// 当 call.Done 说明 结果已经完成了
	replyCall := <- call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	}else {
		fmt.Printf(" async done : %d * %d = %d", args.A, args.B, reply.C)
	}

}

func startServer() {
	s := server.NewServer()
	s.RegisterName("Arith" , new(Arith) , "")
	s.Serve("tcp" , ":8972")
}