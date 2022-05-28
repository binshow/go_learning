package main

import (
  "fmt"
  "sync"
)

//func main() {
//
//  var state int32 = 1
//  fmt.Println(mutexLocked) // 1
//  fmt.Println(1 | mutexLocked) // 1
//  fmt.Println(0 | mutexLocked) // 1
//
//  fmt.Println(2 | mutexLocked) // 3
//  fmt.Println(4 | mutexLocked) // 5   100 | 1 = 101
//  fmt.Println(5 | mutexLocked) // 5   101 | 1 = 101
//
//  // | 表示按位运算，如果 两个数中其中一个为1，则结果该位就是1
//  fmt.Println(state)
//
//}

type Counter struct {
  sync.Mutex
  Count int
}

func main() {
  var c Counter
  c.Lock()
  defer c.Unlock()
  c.Count++
  foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
  c.Lock()
  defer c.Unlock()
  fmt.Println("in foo")
}
