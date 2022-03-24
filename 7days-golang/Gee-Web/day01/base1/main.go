package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//1. 设置路由, 其实就是设置 DefaultServeMux,将对应的URL给对应的方法去处理
	http.HandleFunc("/" , indexHandler)
	http.HandleFunc("/hello" , helloHandler)

	//2. 启动并监听端口， 第二个参数传入的是一个handler
	// ListenAndServe 会一直阻塞，直到发生error
	log.Fatal(http.ListenAndServe(":9999" , nil))

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	// 返回请求的所有头部内容
	for k , v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n" , k , v)
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// 返回请求的路径
	fmt.Fprintf(w , "URL.Path = %q\n" , req.URL.Path)
}
