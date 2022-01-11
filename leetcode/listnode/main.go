package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

	five := &ListNode{5,nil}
	four := &ListNode{4,five}
	three := &ListNode{3,four}
	two := &ListNode{2,three}
	one := &ListNode{1,two}

	node := reverseBetween(one, 2 , 4)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if(head == nil || head.Next == nil) {
		return head
	}
	cur := head
	for i := 0 ; i < k ; i++{
		if(cur == nil) {
			break
		}
		cur = cur.Next
	}
	newHead := reverse(head , cur)
	head.Next = reverseKGroup(cur , k)
	return newHead
}

func reverse(head *ListNode , tail *ListNode) *ListNode{
	var pre *ListNode = nil
	cur := head
	for cur != tail{
		var help *ListNode = cur.Next
		cur.Next = pre
		pre = cur
		cur = help
	}
	return pre
}

func reverseBetween(head *ListNode, l int, r int) *ListNode {
	if head == nil{
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head;

	cur := dummy
	for i := 1 ; i < l ; i++{
		cur = cur.Next
	}

	tail := cur.Next

	a := cur.Next
	cur.Next = nil

	var pre *ListNode = nil
	var help *ListNode = nil

	for i := l ; i <= r ; i++{
		help = a.Next
		a.Next = pre
		pre = a
		a = help
	}

	tail.Next = help
	cur.Next = pre
	return dummy.Next

}

