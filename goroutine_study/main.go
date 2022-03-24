package main

import (
	"fmt"
	"runtime"
	"time"
)

// 测试 goroutine 切换开销
// 两百万次goroutine切换，耗时九千万纳秒
// 一次切换耗时 = 9000 / 200 = 45纳秒 左右
func main() {
	// 设置一个CPU运行
	runtime.GOMAXPROCS(1)
	before := time.Now().UnixNano()
	fmt.Println(before)
	go cal()
	for i := 0; i < 1000000; i++ {
		runtime.Gosched()
	}
	after := time.Now().UnixNano()
	fmt.Println(after - before) // 89121000 纳秒
}

func cal() {
	for i := 0; i < 1000000; i++ {
		runtime.Gosched()
	}
}
