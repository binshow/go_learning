package main

import (
	"fmt"
)

type myInt int
const n = 13 // 无类型常量，只有自己的默认类型，根据初值来决定。当前默认类型为int

const (
	Apple , Banana = 11 , 22
	StrawBerry , Grape
	Pear , Watermelon

)

func main()  {
	var a myInt = 5
	fmt.Println(a + n)

	var b int64 = 3
	fmt.Println(b + n)
}
