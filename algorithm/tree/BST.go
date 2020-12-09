package tree

import (
	"fmt"
	"math"
)

//二叉搜索树特征
/**
 *1、对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。
 *2、对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST。
 */
//基于这个特征，就意味着二叉搜索树的左子树比右子树小，以及其中序遍历是生序序列等可利用之处解决相关问题

//二叉搜索树上的第k小元素
func kthSmallest(root *TreeNode, k int) int {
	res2, rank = 0, 0
	traverse(root, k)
	return res2
}

var res2, rank = 0, 0

//二叉树的中序遍历
func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	rank++
	if rank == k {
		res2 = root.Val
		return
	}
	traverse(root.Right, k)
}

func convertBST(root *TreeNode) *TreeNode {
	traverse2(root)
	return root
}

var sum int

//二叉树的中序遍历
func traverse2(root *TreeNode) {
	if root == nil {
		return
	}
	traverse2(root.Right)
	sum += root.Val
	root.Val = sum
	traverse2(root.Left)
}

//二叉搜索树中2点间的最小差值
func getMinimumDifference(root *TreeNode) int {
	traverse3(root)
	return res3
}

var res3 = math.MaxInt64

//二叉树的中序遍历
func traverse3(root *TreeNode) {
	fmt.Println(res3)
	if root == nil {
		return
	}
	traverse3(root.Left)
	if root.Left != nil {
		if root.Val-root.Left.Val < res3 {
			res3 = root.Val - root.Left.Val
		}
	}
	if root.Right != nil {
		if root.Right.Val-root.Val < res3 {
			res3 = root.Right.Val - root.Val
		}
	}
	traverse3(root.Right)
}
