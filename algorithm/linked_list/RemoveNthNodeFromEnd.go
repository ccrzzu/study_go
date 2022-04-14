package linked_list


//删除链表倒数第n个节点
func removeNthNodeFromEnd(head *ListNode, n int) *ListNode {
	//fast, slow := &ListNode{}, &ListNode{}
	fast, slow := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}