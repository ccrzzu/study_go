package linked_list


//判断链表是否有环
func hasCycle(head *ListNode) bool {
	fast, slow := &ListNode{}, &ListNode{}
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

//判断链表是否有环 by hash
func hasCycleByHash(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}