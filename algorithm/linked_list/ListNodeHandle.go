package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
func ReverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre = nil
	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

//是否是回文链表
func ListNodeIsPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//如果为奇数，将slow再向前移动一位
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := ReverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

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

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
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

//删除链表倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := &ListNode{}, &ListNode{}
	fast, slow = head, head
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
