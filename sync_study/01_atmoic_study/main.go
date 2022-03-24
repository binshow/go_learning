package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	atomicAdd()

	var addr int64
	addr = atomic.AddInt64(&addr , 1) // 原子+1操作
	fmt.Println("after atomic.AddInt64 , addr = " , addr)

	atomic.StoreInt64(&addr , 10)  // 原子赋值操作
	fmt.Println("after atomic.StoreInt64 , addr = " , addr)

	temp := atomic.LoadInt64(&addr) // 原子加载操作
	fmt.Println("after atomic.LoadInt64 , temp = " , temp)

	var a int64 = 10
	old := atomic.SwapInt64(&a , 20) // 原子交换值
	fmt.Println("old = " , old , "a = " , a)

	var b int64 = 2
	c := atomic.CompareAndSwapInt64(&b , 2  , 1)  // 比较并交换，
	fmt.Println("b = " , b  , "c = " , c)
}

// 比较多goroutine 下普通的 ++ 和 原子++ 的区别
func atomicAdd() {
	var d1  , d2 int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&d1  , 1)
			d2++
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("d1 = " , d1 , " , d2 = "  , d2)
}

// ABA 问题模拟，发现没有解决ABA问题
func ValueDemo() {
	var share uint64 = 1
	wg := sync.WaitGroup{}
	wg.Add(3)
	// 协程1，期望值是1,欲更新的值是2
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapUint64(&share,1,2)
		fmt.Println("goroutine 1",swapped)
	}()
	// 协程2，期望值是1，欲更新的值是2
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
		swapped := atomic.CompareAndSwapUint64(&share,1,2)
		fmt.Println("goroutine 2",swapped)
	}()
	// 协程3，期望值是2，欲更新的值是1
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		swapped := atomic.CompareAndSwapUint64(&share,2,1)
		fmt.Println("goroutine 3",swapped)
	}()
	wg.Wait()
	fmt.Println("main exit")
}
