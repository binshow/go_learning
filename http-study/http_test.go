package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

// -------------------------------------------
// @file          : http_test.go
// @author        : binshow
// @time          : 2022/6/22 9:24 AM
// @description   : 学习http库
// -------------------------------------------

//1. request body can only read once
func readBodyOnce (w http.ResponseWriter, r *http.Request) {
	body , err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w , "read body failed : %v\n" , err)
		return
	}

	// []byte ---> string
	fmt.Fprintf(w , "read the data : %v \n" , string(body))

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
		fmt.Fprintf(w , "get body is nil \n")
	}else{
		fmt.Fprintf(w , "get body not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w , "query is %v\n" , values)
}

// 注意这个 URL 中host不一定有值，唯一肯定有值的就是 Path, 其他的可能都没有值。虽然有很多字段
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(r.URL)
	fmt.Fprintf(w , "query is %v" , string(bytes))
}


// header 可以是 http 预定义的，也可以是自定义的，一般用x开头来表示是自己定义的header
func header(w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(r.Header)
	fmt.Fprintf(w , "header is %v" , string(bytes))
}


// Form 单表要先用 ParseForm 解析
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "before parse form  %v \n" , r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w , "parse form error  %v\n" , r.Form)
	}
	fmt.Fprintf(w , "after parse form  %v\n" , r.Form)
}


func TestHttp(t *testing.T) {
	// only post method request has request body!!
	http.HandleFunc("/body/once" , readBodyOnce)
	http.HandleFunc("/body/getbody" , getBodyIsNil) // localhost:8080/body/getbody

	http.HandleFunc("/url/query" , queryParams)	// localhost:8080/url/query?name=binshow&name=zkd&age=18
	http.HandleFunc("/url/wholeUrl" , wholeUrl)	// localhost:8080/url/query?name=binshow&name=zkd&age=18


	http.HandleFunc("/header" , header)
	http.HandleFunc("/form" , form)	// localhost:8080/form?a=b



	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http server err down")
	}
}









