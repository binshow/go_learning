package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置路由, 其实就是设置 DefaultServeMux
	http.HandleFunc("/" , indexHandler)
	http.HandleFunc("/hello" , helloHandler)

	// 启动并监听端口， 第二个参数传入的是一个handler， 见base2
	http.ListenAndServe(":9999" , nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k , v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n" , k , v)
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w , "URL.Path = %q\n" , req.URL.Path)
}
