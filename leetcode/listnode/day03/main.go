package main

import "container/heap"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}


//[剑指 Offer 22. 链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)
func getKthFromEnd(head *ListNode, k int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy

	for k > 0{
		fast = fast.Next
		k--
	}

	for fast != nil{
		fast = fast.Next
		slow = slow.Next
	}
	return slow

}

//[19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy

	for n > 0{
		fast = fast.Next
		n--
	}

	for fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next

}


//[876. 链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/)
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}