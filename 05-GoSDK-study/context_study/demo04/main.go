package main

import (
	"context"
	"fmt"
	"time"
)

/*
context.WithDeadline()
*/

func main() {

	ctx := context.Background()

	// 接收一个ddl参数，只运行这个上下文运行到某个ddl之前
	cancelCtx, cancelFunc := context.WithDeadline(ctx  , time.Now().Add(time.Second * 5))
	defer cancelFunc()

	// 开启一个 goroutine 带着这个上下文去做某些任务
	go dotask(cancelCtx)
	time.Sleep(time.Second * 6)
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
