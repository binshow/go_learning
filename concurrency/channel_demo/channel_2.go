package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//fmt.Printf("\n=> Select and receive\n")
	//selectRecv()

	fmt.Printf("\n=> Select and drop\n")
	selectDrop()

}

// signalAck shows how to signal an event and wait for an acknowledgement it is done
// It does not only want to guarantee that a signal is received but also want to know when that
// work is done. This is gonna like a double signal.
func signalAck() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	// It blocks on the receive. This Goroutine can no longer move on until we receive a signal.
	ch <- "do this"
	fmt.Println(<-ch)
}

// --------------------------------------
// Unbuffered channel_demo: select and receive
// --------------------------------------

// Select allows a Goroutine to work with multiple channel_demo at a time, including send and receive.
// This can be great when creating an event loop but not good for serializing shared state.


// selectRecv 方法展示了如何使用 select 等待特定的时间来接收某个值
func selectRecv() {
	ch := make(chan string)

	// Wait for some amount of time and perform a send.
	// 启动一个 goroutine ，看做是g1
	go func() {
		// 如果这个任务在 100 毫秒内则可以正常完成，否则就会超时
		//time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		time.Sleep(time.Duration(300) * time.Millisecond)
		ch <- "work"
	}()

	// 主线程等待 工作完成的信号，但不想一直等，而是只等 100 毫秒。如果时间范围内没有接收到信号就不等了
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

	// 上面的代码会有一个问题，就是启动的goroutine没有机会终止掉，也就是产生了所谓的 goroutine 泄漏
	// 因为 ch 是一个无缓存的channel，而g1往这个无缓存的 channel中发送信号，而向无缓存的 channel中发送信号
	// 成功的前提是存在一个 对应的接收者。
	// 但是如果是 发生了超时的情况呢，这个g1就发送不会成功，用于不会结束，展示了 goroutine leak

	// 解决这个问题最简单的办法就是 使用一个大小为1的缓存channel

	// The cleanest way to fix this bug is to use the buffered channel_demo of 1. If this send happens,
	// we don't necessarily have the guarantee. We don't need it. We just need to perform the
	// signal then we can walk away. Therefore, either we get the signal on the other side or we
	// walk away. Even if we walk away, this send can still be completed because there is room in
	// the buffer for that send to happen.
}


// selectSend shows how to use the select statement to attempt a send on a channel_demo for a specific
// amount of time.
func selectSend() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
	// 和 selectRecv存在一样的问题和解决办法
}


// selectDrop shows how to use the select to walk away from a channel_demo operation if it will
// immediately block.
// This is a really important pattern. Imagine a situation where our service is flushed with work
// to do or work is gonna coming. Something upstream is not functioning properly. We can't just
// back up the work. We have to throw it away so we can keep moving on.
// A Denial-of-service attack is a great example. We get a bunch of requests coming to our server.
// If we try to handle every single request, we are gonna implode. We have to handle what we can
// and drop other requests.
// Using this type of pattern (fanout), we are willing to drop some data. We can use buffer that
// are larger than 1. We have to measure what the buffer should be. It cannot be random.
func selectDrop() {
	ch := make(chan int, 5)

	go func() {
		// We are in the receive loop waiting for data to work on.
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	// This will send the work to the channel_demo.
	// If the buffer fills up, which means it blocks, the default case comes in and drop things.
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("send work", i)
		default:
			fmt.Println("drop", i)
		}
	}

	close(ch)
}