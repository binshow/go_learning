package main

import (
	"runtime"
	"time"
)

func main() {
	println(runtime.NumCPU())
	go func() {
		time.Sleep(time.Second * 2)
	}()
	println(runtime.NumGoroutine())
}
