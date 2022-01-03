package main

import (
	"log"
	"time"
)

// 2. 可以自定义日志的抬头信息
func init() {
	log.SetFlags(log.Ldate | log.Llongfile)
	//log.SetFlags(log.Ldate|log.Ltime | log.LUTC) //如果设置了LUTC的话，就会把输出的日期时间转为0时区的日期时间显示。

	// 自定义设置日志前缀
	log.SetPrefix("[binshow]") //[binshow]2021/12/18 /Users/shengbinbin/GolandProjects/go_learning/chapter10_log/demo01/main.go:19: this is a log

}

// 演示 log 包的使用
func main() {

	// 1. 默认输出是带时间戳的
	log.Println("this is a log") //2021/12/18 00:08:09 this is a log
	log.Printf("this is also a log: %v\n", "test")

	//3. Fatal 表示程序遇到了致命的错误，先打印日志，再退出
	log.Fatal("this is a fatal log")

	// log.Panic() 同理
	time.Sleep(time.Second * 3) // 不生效了

}
