package main

import (
	"fmt"
	"net"
)

// 实现一个简易的server，监听端口 8888

// handlerConn 处理一个网络连接
func handlerConn(c net.Conn) {
	defer c.Close()
	for  {
		// read from the connection
		// ... do something
		// write to the connection

	}
}

func main() {
	l , err := net.Listen("tcp" , ":8888")
	if err != nil {
		fmt.Println("listen error : " , err)
		return
	}

	// 死循环等待客户端连接
	for {
		c , err := l.Accept()
		if err != nil{
			fmt.Println("accept error: " , err)
			break
		}
		// start a new goroutine to handle the new connection
		go handlerConn(c)
	}
}
