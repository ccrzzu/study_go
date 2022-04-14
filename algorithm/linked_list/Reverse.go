package linked_list


//反转链表 或者称倒排
func ReverseList(head *ListNode) *ListNode {
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


/**
 *将给出的链表中的节点每\ k k 个一组翻转，返回翻转后的链表
 *如果链表中的节点数不是\ k k 的倍数，将最后剩下的节点保持原样
 *你不能更改节点中的值，只能更改节点本身。
 *要求空间复杂度 \ O(1) O(1)
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseListReturnTailAndHead(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverseListReturnTailAndHead(head, tail *ListNode) (*ListNode, *ListNode) {
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

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	prev := &ListNode{0, nil} //哑节点，用于返回反转完成的链表
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
			left, right = reverseListReturnTailAndHead(left, right)
			pre.Next = left
		}
		pre = pre.Next
		left = left.Next
		right = right.Next
		count++
	}
	return prev.Next
}
