package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 配置信息
type Config struct {
	NodeName  string
	Addr 	  string
	Count     int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "北京",
		Addr:     "10.77.95.27",
		Count:    rand.Int31(),
	}
}

func main() {
	//1. 使用 atomic.Value 来存取 config
	var config atomic.Value
	config.Store(loadNewConfig())

	//2. 使用 cond
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for true {
			time.Sleep(time.Duration(5 + rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() // 等待配置变更，唤醒别的goroutine去读取新的配置
		}
	}()

	go func() {
		for true {
			cond.L.Lock()
			cond.Wait()		// 等待变更信号
			c := config.Load().(Config) // 读取新的配置
			fmt.Printf("new config: %+v\n" , c)
			cond.L.Unlock()
		}
	}()


	// 阻塞在这
	select {}

}
