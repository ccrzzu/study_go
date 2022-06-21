package linked_list

/**
206
反转链表
*/
func ReverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func ReverseLinkedList2(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}


/**
92
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
注意：需要反转的入参m,n是指第几个节点，在循环中的index范围要注意
*/

//第一种解法，将需要反转的链表区间反转后，再跟两头接上，反转的部分使用反转链表算法即可
//第 1 步：先将待反转的区域反转；
//第 2 步：把 pre 的 next 指针指向反转以后的链表头节点，把反转以后的链表的尾节点的 next 指针指向 succ。
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	//找到要截取的链表的前一个节点
	pre := newHead
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	//找到要截取的链表的最后一个节点
	rightNode := pre
	for j := 0; j < n-m+1; j++ {
		rightNode = rightNode.Next
	}
	// 要截取的链表的最后一个节点的下一个节点
	succ := rightNode.Next
	leftNode := pre.Next

	//反转前注意切断截取的链表的头尾
	pre.Next = nil
	rightNode.Next = nil
	ReverseLinkedList(leftNode)

	//再将反转后的链表
	pre.Next = rightNode
	leftNode.Next = succ

	return newHead.Next
}

//第二种解法 ： 头插法
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

/**
25
*将给出的链表中的节点每k 个一组翻转，返回翻转后的链表
*如果链表中的节点数不是k 的倍数，将最后剩下的节点保持原样
*你不能更改节点中的值，只能更改节点本身。
解题思路：
这是Swap里第24题的加强版，24 是两两相邻的元素，翻转链表。
而 problem 25 要求的是 k 个相邻的元素，翻转链表，problem 相当于是 k = 2 的特殊情况。
*/

//解法一
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				//代表链表中本轮循环的结点数不足k个，则按原样返回即可
				return dummyHead.Next
			}
		}
		next := tail.Next
		head, tail = reverseListReturnNewHeadAndNewTail(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummyHead.Next
}

//工具方法，将指定头和尾的节点反转后返回新的头和尾节点
func reverseListReturnNewHeadAndNewTail(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		temp := cur
		cur = cur.Next
		temp.Next = prev
		prev = temp
	}
	return prev, head
}

//解法二
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	//哑节点，用于返回反转完成的链表
	prev := &ListNode{0, nil}
	prev.Next = head
	pre := prev               //前置节点
	left, right := head, head //分别指向待翻转头尾
	//初始化尾指针
	for i := 0; i < k-1; i++ {
		right = right.Next
	}
	count := 0 //记录步长
	for right != nil {
		//只有步长为k的倍数时进行翻转，可以解决剩余节点不足k的情况
		if count%k == 0 {
			left, right = reverseListReturnNewHeadAndNewTail(left, right)
			pre.Next = left
		}
		pre = pre.Next
		left = left.Next
		right = right.Next
		count++
	}
	return prev.Next
}

//解法三
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverseReturnNewFirst(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

// 工具类 将指定first和last+1节点之间的first到last所有节点反转后返回新的first节点
func reverseReturnNewFirst(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
