package linked_list

/**
141
判断链表是否有环
给 2 个指针，一个指针是另外一个指针的下一个指针。
快指针一次走 2 格，慢指针一次走 1 格。
如果存在环，那么快指针一定会经过若干圈之后追上慢的指针。
*/
func HasCycle(head *ListNode) bool {
	//fast, slow := &ListNode{}, &ListNode{}
	fast := head
	slow := head
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
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}

/**
 142
 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 注意有可能fast指针走了环的很多圈才和slow相遇.
 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 
 结题思路：
 这道题是第 141 题的加强版。在判断是否有环的基础上，还需要输出环的第一个点。
 分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。
 令链表 head 到环的第一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，从环中相遇点回到环的第一个点需要 x3 步。
 那么环的总长度是 x2 + x3 步。
 fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：
 fast 的 t = (x1 + x2 + x3 + x2) / 2
 slow 的 t = (x1 + x2) / 1
 那么得出：
 x1 + x2 + x3 + x2 = 2 * (x1 + x2)
 所以 x1 = x3
 所以 2 个指针相遇以后，如果 fast 继续往前走，slow 指针回到起点 head，
 两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果。
*/
func DetectCycle(head *ListNode) *ListNode {
	//fast, slow := &ListNode{}, &ListNode{}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}
	//上一步的结束有可能不是fast==slow导致的结束，
	//有可能是循环条件的结束，所以要有这一步的校验
	if fast == nil || fast.Next == nil {
		return nil
	}

	//所以 2 个指针相遇以后，如果 fast 继续往前走，slow 指针回到起点 head，
 	//两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果。
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
