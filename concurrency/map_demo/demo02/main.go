package main

import (
	"fmt"
	"time"
)

func main() {
	var m map[int]int
	//m[100] = 100 // panic，map没有初始化

	m = make(map[int]int) // 初始化map
	fmt.Println(m[10]) 	  // 从map中获取不存在的key不会panic，而是得到0值

	ms := make(map[string]string)
	fmt.Println(ms["aa"]) //空

	var c Counter
	c.Website = "baidu.com"
	c.PageCounters["/"]++  // 写在结构体中的map，使用的时候很容易忘记初始化，一定要记得初始化！！！

}

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}




