package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	var nums  = []int{1,2,3,4}
	reverseArray(nums)

	for _, num := range nums {
		fmt.Println(num)
	}



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
