package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}


//[148. 排序链表](https://leetcode-cn.com/problems/sort-list/)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	r := mergeSort(slow.Next)
	slow.Next = nil
	l := mergeSort(head)
	return merge(l , r)
}

func merge(l *ListNode , r *ListNode) *ListNode{
	if l == nil{
		return r
	}
	if r == nil{
		return l
	}

	dummy := &ListNode{}
	cur := dummy
	for l != nil && r != nil{
		if l.Val < r.Val{
			cur.Next = l
			l = l.Next
		}else{
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}

	if l == nil{
		cur.Next = r
	}
	if r == nil{
		cur.Next = l
	}
	return dummy.Next
}


//[147. 对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/)


//[83. 删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	for cur.Next != nil{
		if cur.Val == cur.Next.Val{
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

//[82. 删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)-1
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	fast := head
	slow := dummy
	for fast != nil{
		if fast.Next != nil && fast.Val == fast.Next.Val{
			//1233445
			for fast.Next != nil && fast.Val == fast.Next.Val{
				fast = fast.Next
			}
			fast = fast.Next
			slow.Next = fast
		}else{
			fast = fast.Next
			slow = slow.Next
		}
	}
	return dummy.Next

}


//[138. 复制带随机指针的链表](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)
type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil{
		return head
	}
	//1. 先复制节点
	cur := head
	for cur != nil{
		tem := &Node{}
		tem.Val = cur.Val
		tem.Next = cur.Next
		cur.Next = tem
		cur = cur.Next.Next
	}

	//2. 复制节点的random指针
	cur = head
	for cur != nil{
		if cur.Random != nil{
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	//3. 分离两个链表
	res := head.Next
	cur = head
	for cur.Next != nil{
		help := cur.Next
		// cur help
		cur.Next = help.Next
		cur = help
	}
	return res

}