package main

import (
	"context"
	"fmt"
	"time"
)

/*

context.WithTimeout()

*/

func main() {

	ctx := context.Background()

	// 接收一个超时参数，3秒钟之后任务还没完成就结束这个上下文
	cancelCtx, cancelFunc := context.WithTimeout(ctx  , time.Second * 3)
	defer cancelFunc()

	// 开启一个 goroutine 带着这个上下文去做某些任务
	go dotask(cancelCtx)

	time.Sleep(time.Second * 4)

	  // 当该函数被执行时，ctx.Done() 的这个Channel就会被Close
	time.Sleep(time.Second * 1)
}

func dotask(ctx context.Context) {
	i := 1
	for  {
		select {
		case <- ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err()) // context deadline exceeded
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
