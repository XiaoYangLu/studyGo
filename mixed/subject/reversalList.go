package subject

type ListNode struct {
	Val  string
	Next *ListNode
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
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
		head, head.Next, newHead = head.Next, newHead, head
	}
	return newHead
}
