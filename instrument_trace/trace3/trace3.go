package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 使用 Trace + defer 来实现函数调用链
// 实现函数名自动获取
// 增加goroutine标识
// 实现层次感

var goroutineSpace = []byte("goroutine ")
var mu sync.Mutex
var m = make(map[uint64]int)

func curGoroutineID() uint64 {
	b := make([]byte , 64)
	b = b[:runtime.Stack(b , false)]
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "  "
	}
	fmt.Printf("g[%05d]:%s%s%s\n" ,id , indents,  arrow ,  name)
}

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock()
	indents := m[gid]  // 获取当前gid对应的缩进层次
	m[gid] = indents + 1 // 缩进层次 + 1 后存入map
	mu.Unlock()
	printTrace(gid , name , "->" , indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid , name  , "<-" , indents)
	}
}

func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()
}

func B2() {
	defer Trace()()
	C2()
}

func C2() {
	defer Trace()()
	D()
}



func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()
	A1()
	wg.Wait()
}
