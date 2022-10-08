package main

import (
	"context"
	"fmt"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/8/1 10:14 AM
// @description   : context 包的用法
// -------------------------------------------

func testContextTimeout() {

	// 设置 context 的过期时间
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 设置 goroutine 的处理时间
	go handle(ctx , 5000 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main" , ctx.Err())
	}

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle"  , ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}


// 测试 取消信号的上下文
func testContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go handle(ctx , time.Second * 1)

	time.Sleep(time.Millisecond * 5000)
	cancel()
}


func main() {
	//testContextTimeout()
	testContextCancel()
}