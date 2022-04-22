package linked_list

/**
链表的中间节点（没有环）
如果链表长度是奇数，输出中间结点是中间结点。如果链表长度是双数，输出中间结点是中位数后面的那个结点。
思路：快慢指针，快指针走到头，慢指针就是中间节点
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func middleNode2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//然后判断列表是奇数还是偶数，以下是错误逻辑，无需判断，仔细想想，多此一举
	// length := 0
	// cur := head
	// for cur != nil {
	// 	length++
	// 	cur = cur.Next
	// }
	// if length%2 == 0 {
	// 	return slow.Next
	// }
	return slow
}
