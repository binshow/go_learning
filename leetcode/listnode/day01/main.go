package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}


//[206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	var help *ListNode
	cur := head
	for cur != nil {
		help = cur.Next
		cur.Next = pre
		pre = cur
		cur = help
	}
	return pre
}

//[25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		 return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	for i := 0 ; i < k ; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	newHead := reverse(head , cur)
	head.Next = reverseKGroup(cur , k)
	return newHead

}

func reverse(head , tail *ListNode) *ListNode {
	var pre *ListNode
	var help *ListNode
	cur := head
	for cur != tail {
		help = cur.Next
		cur.Next = pre
		pre = cur
		cur = help
	}
	return pre

}


//[92. 反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)
func reverseBetween(head *ListNode, l int, r int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for i := 1; i < l; i++ {
		cur = cur.Next
	}

	tail := cur.Next

	a := cur.Next
	cur.Next = nil

	var pre , help *ListNode
	for i := l; i <= r; i++ {
		 help = a.Next
		 a.Next = pre
		 pre = a
		 a = help
	}
	tail.Next = help
	cur.Next = pre
	return dummy.Next


}

//[24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil{
		a := cur.Next
		b := a.Next
		c := b.Next
		// dummy a b c
		cur.Next = b
		b.Next = a
		a.Next = c
		cur = a
	}
	return dummy.Next
}

//[234. 回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		 return true
	}
	// 找中点
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}

	a := slow.Next
	slow.Next = nil

	b := reverseList(a)
	c := head
	for b != nil && c != nil {
		if b.Val != c.Val {
			return false
		}
		b = b.Next
		c = c.Next
	}
	return true
}