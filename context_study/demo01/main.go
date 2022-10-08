package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

const timeout = 3

func main() {
	DoHttpHandler()

	runtime.GOMAXPROCS(1)
	fmt.Println("hello")
}

func DoHttpHandler() {
	// 创建一个超时时间为3秒的上下文
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancelFunc()
	handler(ctx , cancelFunc) // 执行对应逻辑, handler 函数最多只能执行 3秒
}

func handler(ctx context.Context, cancelFunc context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep( 1 * time.Second)
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			cancelFunc() // 手动结束
		}
	}
}





