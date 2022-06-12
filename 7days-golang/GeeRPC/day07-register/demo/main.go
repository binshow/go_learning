package main

import (
	geerpc "go_learning/7days-golang/GeeRPC/day03-service/client"
	"log"
	"sync"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/10 12:08 AM
// @description   :  Client
// -------------------------------------------


type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func main() {
	client, _ := geerpc.Dial("tcp","61274")
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		args := &Args{Num1: 1, Num2: 2 * 2}
		var reply int
		if err := client.Call("Foo.Sum", args, &reply); err != nil {
			log.Fatal("call Foo.Sum error:", err)
		}
		log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
	}()
	//}
	wg.Wait()
}