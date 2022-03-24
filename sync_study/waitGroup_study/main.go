package main

import (
	"fmt"
	"sync"
)

func main() {
	//wg()
	errwg1()
}

//1. 用waitgroup 来控制10个goroutine，并行输出 0~9
func wg() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 注意起 goroutine 的时候需要把参数传递进去
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait() //等待所有的goroutine来完成
	fmt.Println("all goroutine done")
}


//2. waitGroup 不能作为参数传递，是一个值拷贝 , 报错
func errwg1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 注意起 goroutine 的时候需要把参数传递进去
		go func(i int , wg sync.WaitGroup) {
			fmt.Println(i)
			defer wg.Done()
		}(i , wg)
	}
	wg.Wait() //等待所有的goroutine来完成
	fmt.Println("all goroutine done")
}


//2. waitGroup 的 Add 需要在起goroutine之前执行
// 注意 add放在goroutine外，Done放在goroutine内，逻辑复杂时使用defer保证调用！！！
func errwg2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 这样可能导致 for 循环执行完了，goroutine还没有开始跑
		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait() //等待所有的goroutine来完成
	fmt.Println("all goroutine done")
}