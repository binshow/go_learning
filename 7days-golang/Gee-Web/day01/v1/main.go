package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/25 11:25 AM
// @description   : http 原生库的使用
// -------------------------------------------

func main() {
	//1. 设置路由, 其实就是设置 DefaultServeMux,将对应的URL给对应的方法去处理
	http.HandleFunc("/path" , requestUrlPath)
	http.HandleFunc("/header" , requestHeader)
	http.HandleFunc("/body" , readBodyOnce)
	http.HandleFunc("/getbody" , getBodyIsNil)
	http.HandleFunc("/queryParams" , queryParams)
	http.HandleFunc("/wholeUrl" , wholeUrl)
	http.HandleFunc("/form" , form)

	//2. 启动并监听端口， 第二个参数传入的是一个handler,
	//	 handler is a interface,如果传入的是实现了handler的实例，那么所有的http请求都会交给这个实例去处理
	//3. ListenAndServe 会一直阻塞，直到发生error
	log.Fatal(http.ListenAndServe(":9999" , nil))
}

// 返回请求的所有头部内容
func requestHeader(w http.ResponseWriter, req *http.Request) {
	for k , v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n" , k , v)
	}
}

// 返回请求的路径
func requestUrlPath(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w , "URL.Path = %q\n" , req.URL.Path)
}


//1. request body can only read once
func readBodyOnce (w http.ResponseWriter, r *http.Request) {
	body , err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w , "read body failed\n: %v" , err)
		return
	}

	// []byte ---> string
	fmt.Fprintf(w , "read the data : %v\n" , string(body))

	// read again , can not read anything and no error!!!
	body , err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w , "read body failed : %v\n" , err)
		return
	}

	fmt.Fprintf(w , "read the data one more time : [%s] and read data length : %v\n" , string(body) , len(body))
}

// 2. request has GetBody ，但是在原生的request中这个是nil
func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprintf(w , "get body is nil")
	}else{
		fmt.Fprintf(w , "get body not nil")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w , "query is %v" , values)
}

// 注意这个 URL 中host不一定有值，唯一肯定有值的就是 Path, 其他的可能都没有值。虽然有很多字段
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(r.URL)
	fmt.Fprintf(w , "query is %v" , string(bytes))
}



// Form 单表要先用 ParseForm 解析
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "before parse form %v\n" , r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w , "parse form error  %v" , r.Form)
	}
	fmt.Fprintf(w , "after parse form  %v\n" , r.Form)
}