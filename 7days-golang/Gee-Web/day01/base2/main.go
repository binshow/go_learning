package main

import (
	"fmt"
	"log"
	"net/http"
)

//Engine 重写了ServeHTTP方法，实现http.Handler接口
type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}

func main() {
	engine := new(Engine)
	// 传入了 engine，我们就不用DefaultServeMux了
	// 所有的http请求就转向了我们自己的逻辑处理
	log.Fatal(http.ListenAndServe(":9999" , engine))
}
