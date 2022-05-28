package main

import (
	"sync/atomic"
)


// Mutex 第二版：给新人一点机会
type Mutex struct {
	state int32
	sema  uint32
}

// state 是一个复合型字段，一个字段表示多个意义：
const (
	mutexLocked = 1 << iota // 第一位（最小的一位）来表示这个锁是否被持有
	mutexWoken				// 第二位代表是否有唤醒的 goroutine
	mutexWaiterShift = iota // 剩下的代表等待这个锁的 goroutine数
)


func (m *Mutex) Lock() {
	// 目前没有goroutine 持有这个锁， 幸运case，能够直接获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	// 说明这个锁已经被其他的goroutine所持有了

	awoke := false
	// for 循环不断尝试去加锁
	for {

		// 设置新状态，有三个步骤
		old := m.state
		new := old | mutexLocked // 新状态加锁

		// & 运算，只有两个二进制位都是1时，结果为才是1，否则为0
		// old & mutexLocked != 0 说明有等待者了
		if old & mutexLocked != 0 {
			new = old + 1 << mutexWaiterShift //等待者数量加一
		}

		if awoke { // goroutine是被唤醒的
			new &^= mutexWoken // 清除 唤醒标志
		}

		// 通过 CAS 来设置新状态 , 如果cas返回成功说明 抢夺锁的操作成功了，但并不意味着拿到锁了，只是抢锁的操作成功了
		if atomic.CompareAndSwapInt32(&m.state, old, new) {

			// 锁原状态没有加锁，那么这次确实是拿到锁了，直接返回
			if old & mutexLocked == 0 {
				break
			}

			// 锁原来的状态就是被其他的goroutine锁住的，那么这一轮抢夺的操作仅仅就是
			//  清除了唤醒标志 或者 增加了一个 waiter 而已

			// 当前 goroutine 获取 信号量，阻塞等待一段时间
			//todo runtime.Semacquire(&m.sema)

			// 获取到了信号量了，说明当前goroutine是被 唤醒的，设置唤醒标记
			awoke = true
		}
	}
}

func (m *Mutex) Unlock() {
	// Fast path: drop lock bit.

	new := atomic.AddInt32(&m.state, -mutexLocked)  // 去掉锁标志
	if (new + mutexLocked) & mutexLocked == 0 { 	// 本来就没有加锁，解锁就会有问题
		panic("sync: unlock of unlocked mutex")
	}

	// 因为可能有一些等待锁的 goroutine，需要通过信号量的方式来唤醒

	old := new
	for {
		// 没有等待者，或者有唤醒的waiter，或者锁原来已加锁
		// old >> mutexWaiterShift == 0 ：说明没有其他的waiter，对这个锁的竞争只有一个
		// old&(mutexLocked|mutexWoken) != 0 ： 说明有唤醒的goroutine 或者是又已经被别人加了锁了
		if old >> mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 {
			return
		}

		// 代码到这里，说明有 等待者，而且没有已经被唤醒的waiter

		// 将等待者数量-1 ，准备唤醒goroutine，并设置唤醒标志
		new = (old - 1<<mutexWaiterShift) | mutexWoken
		if atomic.CompareAndSwapInt32(&m.state, old, new) {

			//todo runtime.Semrelease(&m.sema) 实际唤醒了一个等待者
			return
		}
		old = m.state
	}
}


