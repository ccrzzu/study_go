package linked_list

/**
148
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	newHead := slow.Next
	slow.Next = nil

	l1 := sortList(head)
	l2 := sortList(newHead)

	return mergeTwoLists(l1, l2)
}
