package main

import (
	"MyProject/algorithm/array"
	"MyProject/algorithm/stack"
	"MyProject/algorithm/tree"
	"fmt"
	"testing"
)

func TestIsBalancedTree(t *testing.T) {
	root := &tree.TreeNode{
		Val: 3,
	}
	node9 := &tree.TreeNode{
		Val: 9,
	}
	node20 := &tree.TreeNode{
		Val: 20,
	}
	node15 := &tree.TreeNode{
		Val: 15,
	}
	node7 := &tree.TreeNode{
		Val: 7,
	}
	node20.Left = node15
	node20.Right = node7
	root.Left = node9
	root.Right = node20
	//fmt.Println(tree.MaxDepth(root.Left),tree.MaxDepth(root.Right))
	fmt.Println(tree.IsBalanced(root))
}

func TestArrayMerge(t *testing.T) {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3
	array.MergeB2A(A, m, B, n)
}

func TestIsPopOrder(t *testing.T) {
	pushArr := []int{1, 2, 3, 4, 5}
	popArr := []int{4, 5, 3, 2, 1}
	fmt.Println(stack.IsPopOrder(pushArr, popArr))
}

func TestSlice(t *testing.T) {
	var dp1 []int
	dp2 := make([]int, 5, 10)
	fmt.Println(dp1 == nil, dp2 == nil)
	dp2[1] = 1
	fmt.Println(dp2, cap(dp2))
	dp2 = append(dp2, 1, 2, 3, 4)
	fmt.Println(dp2, len(dp2), cap(dp2))
	dp2[8] = 8
	fmt.Println(dp2, len(dp2), cap(dp2))
	dp2[9] = 9
	fmt.Println(dp2, len(dp2), cap(dp2))
}

func TestKeyWordNew(t *testing.T) {
	var sum *int
	sum = new(int) //分配空间
	fmt.Println(sum)
	var a []int
	var b = new([]int)
	fmt.Println(a == nil, b == nil, len(a), len(*b))
	for range a {

	}
	for range *b {

	}
}
