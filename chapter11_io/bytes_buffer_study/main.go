package main

import (
	"bytes"
	"fmt"
)

/**
	测试 buffer
 */

func main() {

	//1. 声明一个 buffer
	var buffer bytes.Buffer

	//2. 往 buffer 里面写数据
	n, err := buffer.WriteString("this is a  test")
	fmt.Println(n , err) // 15 <nil>
	fmt.Println(buffer.Len() , buffer.Cap()) // 15 ,64


	//3. 从 buffer 里面读数据
	s := make([]byte , 1000)
	n, err = buffer.Read(s)
	fmt.Println(n , err) // 15 nil
	fmt.Println(buffer.Len() , buffer.Cap()) //0 64
	fmt.Println(string(s)) //this is a  test

}
