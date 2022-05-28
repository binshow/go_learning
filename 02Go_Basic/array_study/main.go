package main

import (
	"crypto/sha256"
	"fmt"
)

func foo(nums [5]int) {}

func main() {
	//var arr1 [5]int
	//var arr2 [6]int
	//var arr3 [5]string
	//foo(arr1)
	//foo(arr2) // Cannot use 'arr2' (type [6]int) as the type [5]int
	//foo(arr3) // Cannot use 'arr3' (type [5]string) as the type [5]int
}

func Sum256Demo() {
	// Sum256 对一个任意的字节slice类型的数据生成一个对应的消息摘要
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1 = %v \n, c2 = %v \n" , c1 , c2)
	fmt.Printf("%T" , c1 ) //[32]uint8
}

func reverseArray(nums []int) {
	for index, val := range nums {
		nums[index] = -val
	}
}
