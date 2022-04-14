package linked_list

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