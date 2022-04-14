package linked_list


//链表的中间节点
func middleNode(head *ListNode) *ListNode {
	fast, slow := &ListNode{}, &ListNode{}
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}