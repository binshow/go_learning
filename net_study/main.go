package main

import (
	"fmt"
	"net"
	"net/http"
)

// go net 包基本原理演示：看起来就是一个同步模型，Accept、Read 和 Write 都会将当前协程给“阻塞”掉。
// 比如 Read 函数这里，如果服务器调用时客户端数据还没有到达，
// 那么 Read 是不带返回的，会将当前的协程 park 住。直到有了数据 Read 才会返回，处理协程继续执行。
// 被阻塞的 goroutine 是如何唤醒的呢？ 是通过一个周期性的后台监控 goroutine：sysmon
func main() {
	//1. 根据协议名称和地址创建一个 Listener
	listener, _ := net.Listen("tcp", "127.0.0.1:8089")
	for {
		//2. 调用 listener 的 Accept方法等待客户端的连接进来
		// Accept 会阻塞
		conn, _ := listener.Accept()

		//3. 如果有连接进来，就开启一个 goroutine 去处理这个连接
		go process(conn)
	}

	http.ListenAndServe()

}

func process(conn net.Conn) {
	//结束时关闭连接
	defer conn.Close()

	//读取连接上的数据
	var buf [1024]byte

	len, err := conn.Read(buf[:]) // Read 会阻塞

	//发送数据
	_, err = conn.Write([]byte("I am server!")) // Write 会阻塞
	fmt.Println(len, err)
}

