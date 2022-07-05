package main

import (
	"errors"
	"fmt"
	"os"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/25 3:50 PM
// @description   : error vs exception
// -------------------------------------------


///Users/shengbinbin/sdk/go1.17.2/src/builtin/builtin.go:261


//1. error 就是go语言内置的一个 interface
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
//	type error interface {
//		Error() string
//	}

//2. 比较两个相同描述的error的场景，不能用字符串相等来判断两个error！！！
// 可能比较地址更适用一点，因为项目中两个地方爆出了相同的err，其实是两个error
func diffTwoError() {
	err1 := errors.New("my error")
	fmt.Println(err1)

	err2 := errors.New("my error")
	fmt.Println(err1 == err2) // false , 因为比较的是指针地址
}

//3. 适用自定义的 errorString type
func useErrorString() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")	// 这里不输出，因为内存地址不相同
	}
}

func main() {

	useErrorString()
	os.PathError{}


}