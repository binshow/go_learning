package main

type ListNode struct {
	Val int
	Next *ListNode
}

//[86. 分隔链表](https://leetcode-cn.com/problems/partition-list/)
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	small := dummy1
	big := dummy2;

	cur := head
	for cur != nil{
		if cur.Val < x{
			small.Next = cur
			cur = cur.Next
			small = small.Next
		}else{
			big.Next = cur
			cur = cur.Next
			big = big.Next
		}
	}

	small.Next = dummy2.Next
	big.Next = nil // 记得置空！！！
	return dummy1.Next

}


//[725. 分隔链表](https://leetcode-cn.com/problems/split-linked-list-in-parts/)
func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode , k)
	if head == nil {
		return res
	}

	len := 0
	cur := head
	for cur != nil{
		cur = cur.Next
		len++
	}

	size := len / k  // 分成size段
	mod := len % k   // 前mod段多一个

	cur = head
	for i := 0 ; i < k && cur != nil ; i++{
		res[i] = cur
		var curSize int
		if mod > 0{
			curSize = size + 1
			mod--
		}else{
			curSize = size
		}

		for j:= 0 ; j < curSize-1 ; j++{
			cur = cur.Next
		}

		help := cur.Next
		cur.Next = nil
		cur = help
	}

	return res

}


//[2. 两数相加](https://leetcode-cn.com/problems/add-two-numbers/)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil{
		var a , b int
		if l1 != nil{
			a = l1.Val
		}
		if l2 != nil{
			b = l2.Val
		}
		sum := a + b + carry
		carry = sum / 10
		cur.Next = &ListNode{sum % 10 , nil}
		cur = cur.Next

		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
	}

	if carry == 1{
		cur.Next = &ListNode{1 , nil}
	}
	return dummy.Next


}



//[146. LRU 缓存机制](https://leetcode-cn.com/problems/lru-cache/)
type LinkNode struct {
	key ,val  int
	pre ,next *LinkNode
}

type LRUCache struct {
	m map[int]*LinkNode
	cap int
	dummyHead , dummyTail *LinkNode
}

func Constructor(capacity int) LRUCache {
	dummyHead := &LinkNode{0 , 0 , nil ,nil}
	dummyTail := &LinkNode{0 , 0 , nil , nil}
	dummyHead.next = dummyTail
	dummyTail.pre = dummyHead
	return LRUCache{make(map[int]*LinkNode) , capacity , dummyHead, dummyTail}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v , exist := cache[key] ; exist{
		this.MoveToHead(v)
		return v.val
	}else{
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	tail := this.dummyTail
	cache := this.m
	if v , exist := cache[key];exist{
		v.val = value
		this.MoveToHead(v)
	}else{
		curNode := &LinkNode{key , value , nil , nil}
		if len(cache) == this.cap {
			// 删除节点
			delete(cache , tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		this.AddHeadNode(curNode)
		cache[key] = curNode
	}

}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddHeadNode(node)
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddHeadNode(node *LinkNode) {
	head := this.dummyHead
	node.next = head.next
	head.next.pre = node
	head.next = node
	node.pre = head
}


