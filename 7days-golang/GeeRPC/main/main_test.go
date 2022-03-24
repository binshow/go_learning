package main

import (
	"encoding/json"
	"fmt"
	"go_learning/7days-golang/GeeRPC"
	codec2 "go_learning/7days-golang/GeeRPC/codec"
	"log"
	"net"
	"testing"
	"time"
)

func TestCodec(t *testing.T) {
	// 开启服务端
	addr := make(chan string)
	go startServer(addr)


	// in fact, following code is like a simple geerpc client
	// 模拟客户端发送请求
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(GeeRPC.DefaultOption)

	// 拿到新的编码器
	cc := codec2.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec2.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
