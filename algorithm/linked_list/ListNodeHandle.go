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

//判断链表是否有环 by hash
func hasCycleByHash(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil{
		if _,ok := m[head];ok{
			return true
		}
		m[head] = 1
		head = head.Next
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

//合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	result := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return result.Next
}

//删除链表中的某节点
func deleteNode(node *ListNode) {
	if node == nil{
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//删除链表中的某节点,返回头节点
func deleteNodeReturnHead(head *ListNode, val int) *ListNode {
	tmp := head
	for tmp.Next != nil{
		if tmp.Next.Val == val{
			tmp.Val = tmp.Next.Val
			tmp.Next = tmp.Next.Next
			break
		}
		tmp = tmp.Next
	}
	return head
}