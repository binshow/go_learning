package main

import (
	"fmt"
	"go_learning/7days-golang/Gee-Web/day01/base3/gee"
	"net/http"
)

func main() {

	// 使用自定义的Gee框架
	r := gee.New()

	// Get 其实就是往路由表中增加映射
	r.GET("/" , func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w , "URL.Path = %q \n" , r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}