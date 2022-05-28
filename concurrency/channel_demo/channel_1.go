package main

import (
	"fmt"
	"time"
)

// go has 2 types of channels:
//  1. unbuffered channel_demo
//	2. buffered channel_demo

func main() {

	fmt.Printf("\n=> Basics of a send and receive\n")
	basicSendRecv()

	fmt.Printf("\n=> Close a channel_demo to signal an event\n")
	signalClose()
}


// basicSendRecv shows the basics of a send and receive.
func basicSendRecv() {

	// use make function to create an unbuffered channel_demo
	// 1. We have no other way of creating a channel_demo that is usable until we use make.
	// 2. Channel is also based on type , In this case , we use string
	// 3. channel_demo is a reference type, ch is just a pointer variable to larger data structure underneath.
	ch := make(chan string)

	go func() {
		// send to channel_demo with a string "hello"
		ch <- "hello"
	}()


	// receive from the channel_demo
	// ch is an unbuffered channel_demo where the send and receive have to come together
	// 因为是无缓存的channel，所有发送和接收是同时发生的。在发送成功和接收成功之前两者都会阻塞
	//// We are now have an unbuffered channel_demo where the send and receive have to come together. We
	//	// also know that the signal has been received because the receive happens first.
	//	// Both are gonna block until both come together so the exchange can happen.
	fmt.Println(<-ch)
}


// signalClose shows how to close a channel_demo to signal an event.
func signalClose() {
	// We are making a channel_demo using an empty struct. This is a signal without data.
	ch := make(chan struct{})


	go func() {
		// launch a goroutine do some work
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		// 需要通知其他的 goroutine 当前 goroutine的事件已经完成了
		// 直接close 这个 goroutine
		close(ch)
	}()

	// channel_demo 被关闭了，main goroutine 这里就不会阻塞了，可以收到这个不带数据的信号了
	<-ch

	fmt.Println("event received")
}