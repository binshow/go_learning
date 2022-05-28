package main

import (
	"fmt"
	"sync"
	"time"
)
// unbuffered channel_demo 演示案例二
// 使用 unbuffered channel_demo 来模拟4个goroutine的接力赛

// 赛道上有4个运动员，同一时间内只能有一个运动员在跑路，上一个运行员跑完了下一个才能开始跑

// 并发控制
var wg sync.WaitGroup

func main() {

	track := make(chan int)

	// 加1 是为了保证最后一个运动员跑完之后程序结束
	wg.Add(1)

	// 第一个运动员跑起来了
	go Runner(track)


	// main goroutine shoot the gun ， 比赛开始了
	track <- 1

	wg.Wait()

}

func Runner(track chan int) {

	// 最大的接力棒交换次数
	const maxExchanges = 4

	var exchange int

	baton := <- track

	// Start running around the track.
	fmt.Printf("Runner %d Running With Baton\n", baton)

	// New runner to the line. Are we the last runner on the race?
	// If not, we increment the data by 1 to keep track which runner we are on.
	// We will create another Goroutine. It will go immediately into a receive. We are now having a
	// second Groutine on the track, in the receive waiting for the baton. (1)
	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("Runner %d To The Line\n", exchange)
		go Runner(track)
	}

	// Running around the track.
	// 当前 goroutine 跑完所需要的时间
	time.Sleep(100 * time.Millisecond)


	// Is the race over.
	if baton == maxExchanges {
		fmt.Printf("Runner %d Finished, Race Over\n", baton)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\n", baton, exchange)

	// Since we are not the last runner, perform a send so (1) can receive it.
	track <- exchange


}
