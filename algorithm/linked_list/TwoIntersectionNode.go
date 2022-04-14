package linked_list

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
//如果两个链表不存在相交节点，返回 null 。
/**
这道题的思路其实类似链表找环。
给定的 2 个链表的长度如果一样长，都从头往后扫即可。如果不一样长，需要先“拼成”一样长。
把 B 拼接到 A 后面，把 A 拼接到 B 后面。这样 2 个链表的长度都是 A + B。再依次扫描比较 2 个链表的结点是否相同。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	// A拼接B B拼接A 后 两个链表最终一定能走到a==b
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
