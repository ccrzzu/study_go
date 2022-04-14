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