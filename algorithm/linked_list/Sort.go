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

/**
排序奇升偶降链表（字节高频题）
问题描述：给定一个奇数位升序，偶数位降序的链表，将其重新排序。不能使用额外空间。
案例：
输入: 1->8->3->6->5->4->7->2->NULL
输出: 1->2->3->4->5->6->7->8->NULL

思路：核心主要分为以下三步：
1、按照奇偶位置拆分链表（leetcode328）
2、反转偶链表（leetcode206）
3、合并两个有序链表（leetcode21）
ps：在拆分链表时遇到一个坑，一定要 odd.next = null， 否则出现死循环
*/

/**
328
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
*/
func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{Val: 0, Next: nil}
	odd := oddHead
	evenHead := &ListNode{Val: 0, Next: nil}
	even := evenHead

	count := 1
	for head != nil {
		if count%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		count++
	}
	even.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
}
