package main

import "fmt"

func main() {
	//case1 为空链表的情况
	var head *ListNode = nil
	fmt.Println(CheckListNodes(head))

	//case2 正常偶数个的情况，默认是回文链表返回true，随意改个其中某个Node的val值为别的字母即是非回文返回false
	nodeF := &ListNode{Val: "a", Next: nil}
	nodeE := &ListNode{Val: "b", Next: nodeF}
	nodeD := &ListNode{Val: "c", Next: nodeE}
	nodeC := &ListNode{Val: "c", Next: nodeD}
	nodeB := &ListNode{Val: "b", Next: nodeC}
	nodeA := &ListNode{Val: "a", Next: nodeB}
	head = nodeA
	fmt.Println(CheckListNodes(head))

	//case3 正常奇数个的情况，默认是回文链表返回true，除了中间字母，随意改个其他某个Node的val值为别的字母即是非回文返回false
	nodeG := &ListNode{Val: "a", Next: nil}
	nodeF = &ListNode{Val: "b", Next: nodeG}
	nodeE = &ListNode{Val: "c", Next: nodeF}
	nodeD = &ListNode{Val: "d", Next: nodeE}
	nodeC = &ListNode{Val: "c", Next: nodeD}
	nodeB = &ListNode{Val: "b", Next: nodeC}
	nodeA = &ListNode{Val: "a", Next: nodeB}
	head = nodeA
	fmt.Println(CheckListNodes(head))

	//case2 只有一个节点的情况
  	nodeA = &ListNode{Val: "a", Next: nodeB}
  	head = nodeA
  	fmt.Println(CheckListNodes(head))

}

type ListNode struct {
	Val  string
	Next *ListNode
}

func CheckListNodes(head *ListNode) bool {
	//链表区分个数是奇数偶数
	//偶数的话只需要将后半段反转后与前半段元素逐个作对比
	//奇数的话只需要将除了中间位置的元素之外的后半段元素反转后与前半段元素逐个作对比
	// a-->b-->c-->c-->b-->a
	// a-->b-->c-->d-->c-->b-->a
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//fmt.Println(slow.Val, fast.Val)
	}
	if fast != nil{
		slow = slow.Next
	}
	left := head
	right := reverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
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
