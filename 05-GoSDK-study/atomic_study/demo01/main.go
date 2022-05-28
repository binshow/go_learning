package main

import (
	"fmt"
	"sync/atomic"
)

const x int64 = 1 + 1<<33

func main() {

	var i int32 = 3
	if atomic.CompareAndSwapInt32(&i, 3, 5){ // 比较并交换
		fmt.Println(i)
	}

}

func AddxxxDemo() {
	var i int32 = 1
	println(atomic.AddInt32(&i, 1)) // 2
	fmt.Println(i) // 2

	println(atomic.AddInt32(&i, ^int32(0))) // 1 ,通过补码的方式实现原子减法
	fmt.Println(i) // 1
}

func ValueDemo() {

}
