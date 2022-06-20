package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // 需要导入 pprof 的包！！！
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/19 3:33 PM
// @description   :	pprof 性能分析第一次使用
// -------------------------------------------

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}