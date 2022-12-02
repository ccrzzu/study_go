package tree

import (
	"fmt"
	"math"
)

//判断一棵树是否是平衡二叉树
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !IsBalanced(root.Left) || !IsBalanced(root.Right) {
		return false
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if root.Left != nil && root.Right != nil {
		fmt.Printf("node:%v,depth:%v,node:%v,depth:%v\n", root.Left.Val, leftDepth, root.Right.Val, rightDepth)
	}
	if math.Abs(float64(leftDepth)-float64(rightDepth)) > 1 {
		return false
	}
	return true
}