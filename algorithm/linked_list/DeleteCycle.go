package linked_list


//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//注意有可能fast指针走了环的很多圈才和slow相遇
func detectCycle(head *ListNode) *ListNode {
	fast, slow := &ListNode{}, &ListNode{}
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}
	//上一步的结束有可能不是fast==slow，所以要有这一步的校验
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}