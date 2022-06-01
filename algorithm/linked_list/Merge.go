package linked_list

/**
21
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	return result.Next
}

// 通过递归 解法二
// 注意：递归的总时间复杂度 = 每次递归的时间复杂度 * 递归的深度
func mergeTwoListsByDG(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsByDG(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsByDG(l1, l2.Next)
	return l2
}

/**
23 合并k个排序链表
给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。
解题思路：借助分治的思想，把 K 个有序链表两两合并即可。相当于是第 21 题的加强版。
*/
func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := MergeKLists(lists[:num])
	right := MergeKLists(lists[num:])
	return mergeTwoListsByDG(left, right)
}

/**
1699
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cur := list1
	var aPrev, bCur *ListNode
	for i := 0; i <= b; i++ {
		if i == a-1 {
			aPrev = cur
		}
		if i == b {
			bCur = cur
		}
		cur = cur.Next
	}
	aPrev.Next = list2
	cur = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = bCur.Next
	return list1
}
