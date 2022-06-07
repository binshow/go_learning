package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 使用CAS来代替锁实现转账服务

func main() {
	//transferByLock()
	//transferByWaitGroup()

	now := time.Now()
	for i := 0; i < 20; i++ {
		//transferByWaitGroup()
		transferByCAS2()
	}
	fmt.Println(time.Since(now).String())

}

//1. 使用锁来完成转账
func transferByLock() {
	var balance int32 = 0
	done := make(chan bool)
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			balance += 1
			lock.Unlock()
			done <- true
		}()
	}
	time.Sleep(time.Second * 1)  //等待所有转账完成
	fmt.Printf("balance = %d \n" , balance)
}

//2. 使用 waitGroup + Lock 完成转账
func transferByWaitGroup() {
	var balance int32 = 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			balance += 1
			lock.Unlock()
			defer wg.Done()
		}()
	}
	wg.Wait() //等待所有转账完成
	//fmt.Printf("balance = %d \n" , balance)
}


//3. 使用 cas 完成转账
func transferByCAS() {
	var balance int32 = 0
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// 死循环 + CAS
			for  {
				old := atomic.LoadInt32(&balance)
				new := old + 1
				if atomic.CompareAndSwapInt32(&balance, old, new) {
					break
				}
			}
			defer wg.Done()
		}()
	}
	wg.Wait() //等待所有转账完成
	//fmt.Printf("balance = %d \n" , balance)
}


//3. 使用 cas 完成转账 2
func transferByCAS2() {
	var balance int32 = 0
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// 死循环 + CAS
			atomic.AddInt32(&balance , 1)
			defer wg.Done()
		}()
	}
	wg.Wait() //等待所有转账完成
	fmt.Printf("balance = %d \n" , balance)
}