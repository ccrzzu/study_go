package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//注意有可能fast指针走了环的很多圈才和slow相遇
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
	//上一步的结束有可能不是fast==slow，所以要有这一步的校验
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
	if node == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//删除链表中的某节点,返回头节点
func deleteNodeReturnHead(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy.Next
	for second != nil {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}
	return dummy.Next
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//个位在头结点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i, j := l1, l2
	head := new(ListNode)
	tmp := head
	var carry, digit int
	for i != nil || j != nil || carry > 0 {
		digit = carry
		if i != nil {
			digit += i.Val
			i = i.Next
		}
		if j != nil {
			digit += j.Val
			j = j.Next
		}
		if digit >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = digit % 10
	}
	return head.Next
}

//两个链表相加，个位在尾结点，比上难
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	l1, l2 := ReverseList(head1), ReverseList(head2)
	head := new(ListNode)
	tmp := head
	var carry, digit int
	for l1 != nil || l2 != nil || carry > 0 {
		digit = carry
		if l1 != nil {
			digit += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			digit += l2.Val
			l2 = l2.Next
		}
		if digit >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = digit % 10
	}
	return ReverseList(head.Next)
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
	prev := &ListNode{0, nil}//哑节点，用于返回反转完成的链表
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