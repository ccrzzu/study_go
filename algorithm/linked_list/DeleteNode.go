package linked_list

//删除链表中的某节点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//删除链表中的某节点,返回头节点
func deleteNodeReturnHead(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy.Next
	for second != nil {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}
	return dummy.Next
}

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