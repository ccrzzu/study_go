package linked_list

//2 
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//个位在头结点
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

//两个链表相加，个位在尾结点，比上面的难
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