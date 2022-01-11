package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){

}


//[143. 重排链表](https://leetcode-cn.com/problems/reorder-list/)
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil{
		return
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}

	a := head
	b := reverse(slow.Next)
	slow.Next = nil

	for  a != nil && b != nil{
		c := a.Next
		d := b.Next
		// a  c
		// b  d
		a.Next = b
		b.Next = c

		a = c
		b = d
	}
	return
}

func reverse(head *ListNode) *ListNode{
	var pre *ListNode
	cur := head
	for cur != nil{
		help := cur.Next
		cur.Next = pre
		pre = cur
		cur = help
	}
	return pre
}

//[328. 奇偶链表](https://leetcode-cn.com/problems/odd-even-linked-list/)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	a1 := &ListNode{}
	b1 := &ListNode{}

	a1.Next = head
	b1.Next = head.Next

	a := head
	b := head.Next
	for b != nil && b.Next != nil{
		c := b.Next
		// a b  c
		a.Next = c
		b.Next = c.Next

		a = c
		b = c.Next
	}
	a.Next = b1.Next
	return head

}


//[160. 相交链表](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil {
		return headB
	}
	if headB == nil{
		return headA
	}

	l1 := headA
	l2 := headB
	for l1 != l2{
		if l1 != nil{
			l1 = l1.Next
		}else{
			l1 = headB
		}

		if l2 != nil{
			l2 = l2.Next
		}else{
			l2 = headA
		}
	}
	return l1


}


//[61. 旋转链表](https://leetcode-cn.com/problems/rotate-list/)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}

	len := 1
	cur := head
	for cur.Next != nil{
		cur = cur.Next
		len++
	}
	cur.Next = head
	x := len - k % len

	a := cur
	for x > 0{
		a = a.Next
		x--
	}

	res := a.Next
	a.Next = nil
	return res

}


//[141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)
func hasCycle(head *ListNode) bool {
		if head == nil || head.Next == nil{
			return false
		}

		fast := head
		slow := head
		for fast.Next != nil && fast.Next.Next != nil{
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow{
				return true
			}
		}
		return false
}


//[142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			fast = head
			for fast != slow{
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}