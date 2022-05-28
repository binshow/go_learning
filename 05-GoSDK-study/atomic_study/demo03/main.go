package main

import (
	"sync/atomic"
	"unsafe"
)

// 使用atomic来实现 Lock-Free queue

// LKQueue lock-free queue
type LKQueue struct {
	head unsafe.Pointer  // 头节点指针
	tail unsafe.Pointer
}

// node: 节点
type node struct {
	value interface{}
	next unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n , tail: n}
}

// Enqueue 入队：通过CAS操作将一个元素添加到队尾，并且移动尾指针
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v} // 构造节点
	for {
		tail := load(&q.tail)
		next := load(&tail.next)

		if tail == load(&q.tail){ // 尾节点还是尾节点
			if next == nil { // 还没有新数据入队
				if cas(&tail.next, next, n) {  // 将新节点增加到队尾
					cas(&q.tail , tail , n)    // 入队成功了，移动尾指针
					return
				}
			}else { // 已经有新数据了
				cas(&q.tail , tail , next)
			}
		}
	}
}

// Dequeue 出队，没有元素返回nil，出队的时候移除一个节点，并通过CAS操作移动head指针，同时在必要的时候移动尾指针
func (q *LKQueue) Dequeue() interface{}{
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)

		if head == load(&q.head) {
			if head == tail {
				if next == nil {  // 说明是空队列
					return nil
				}
				cas(&q.tail , tail , next) // 只是尾指针还没有调整，尝试调整它指向下一个
			}else {
				v := next.value // 读取出队的数据
				if cas(&q.head, head, next) { // 移动头指针成功
					return v
				}
			}
		}
	}
}

// 将 unsafe.Pointer 原子加载为 node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装cas，避免直接将*node转换成unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p , unsafe.Pointer(old) , unsafe.Pointer(new))
}


func main() {

}
