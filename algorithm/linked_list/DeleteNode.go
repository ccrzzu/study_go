package linked_list

//237 删除链表中的某节点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

/**
83
给定一个已排序的链表的头head ，删除所有重复的元素，使每个元素只出现一次 。
返回 已排序的链表 。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

/**
203
给你一个链表的头节点 head 和一个整数 val ，
请你删除链表中所有满足 Node.Val == val 的节点，并返回 新的头节点 。
*/
func deleteNodeReturnHead(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{0, head}
	pre := newHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return newHead.Next
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
