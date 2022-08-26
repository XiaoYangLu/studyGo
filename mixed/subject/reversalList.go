package subject

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Re() {
	h1 := new(ListNode)
	h1.Val = 1
	h2 := new(ListNode)
	h2.Val = 2
	h3 := new(ListNode)
	h3.Val = 3
	h4 := new(ListNode)
	h4.Val = 4
	h1.Next = h2
	h2.Next = h3
	h3.Next = h4
	pre := reverseList(h1)
	fmt.Println(pre)
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	var newHead *ListNode
	//for head != nil {
	//	node := head.Next
	//	head.Next = newHead
	//	newHead = head
	//	head = node
	//}
	for head != nil {
		newHead, head, head.Next = head, head.Next, newHead
	}
	return newHead
}
