package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "welcome , gohelper!\n")
}

func main() {
	//  用Go搭建一个Web Server
	// http.ListenAndServe(":8080" , greeting) // 编译报错

	// 通过适配器完成函数转型
	http.ListenAndServe(":8080" , http.HandlerFunc(greeting))
}
