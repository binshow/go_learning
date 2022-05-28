package main

import (
	"fmt"
	"io"
	"net/http"
)
// 定义自己的函数类型
type HandlerFunc  func(w http.ResponseWriter , r http.Request)

var (
	// 匿名函数存储在变量中
	myFprintf = func(w io.Writer , format string , a...interface{})(int , error) {
		return fmt.Fprintf(w , format , a...)
	}
)

func main() {
	teardown := setup("binshow")
	defer teardown()
	fmt.Println("do some bussiness stuff")
}

func setup(task string) func() {
	fmt.Println("do some setup stuff for" , task)
	// 函数内创建并通过返回值返回
	return func() {
		println("do some teardown stuff for" , task)
	}
}
