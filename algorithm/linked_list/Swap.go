package linked_list

/**
 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
} 