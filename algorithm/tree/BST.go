package tree

import (
	"math"
)

//二叉搜索树特征
/**
 *1、对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。
 *2、对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST。
 */
//基于这个特征，就意味着二叉搜索树的左子树比右子树小，以及其中序遍历是生序序列等可利用之处解决相关问题

//判断二叉搜索树的有效性
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBSTByDG(root, math.MinInt64, math.MaxInt64)
}
func isBSTByDG(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isBSTByDG(root.Left, min, root.Val) && isBSTByDG(root.Right, root.Val, max)
}

func isValidBST2(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode, *TreeNode) bool
	dfs = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}
		//左子树不能超过最大值
		if max != nil && node.Val >= max.Val {
			return false
		}
		//右子树不能小于最小值
		if min != nil && node.Val <= min.Val {
			return false
		}
		return dfs(node.Left, min, node) && dfs(node.Right, node, max)
	}
	return dfs(root, nil, nil)
}

//二叉搜索树上的第k小元素
func kthSmallest(root *TreeNode, k int) int {
	res, rank := 0, 0
	var traverse func(*TreeNode, int)
	traverse = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		traverse(root.Left, k)
		rank++
		if rank == k {
			res = root.Val
			return
		}
		traverse(root.Right, k)
	}
	traverse(root, k)
	return res
}

//二叉搜索树上的第k大元素
func kthLargest(root *TreeNode, k int) int {
	res, rank := 0, 0
	var traverse func(*TreeNode, int)
	traverse = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		traverse(root.Right, k)
		rank++
		if rank == k {
			res = root.Val
			return
		}
		traverse(root.Left, k)
	}
	traverse(root, k)
	return res
}

//将一棵二叉搜索树改为累积树，右子树及本身累加
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}
	traverse(root)
	return root
}

//二叉搜索树中2点间的最小差值
//思想仍旧是有序序列的没相邻的2个数之间的差的最小值，就是我们要找的最小差值
func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt64, -1
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if pre != -1 && node.Val-pre < res {
			res = node.Val - pre
		}
		pre = node.Val
		traverse(node.Right)
	}
	traverse(root)
	return res
}

//给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
func findTargetByDG(root *TreeNode, k int) bool {
	var res bool
	m := map[int]bool{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if m[k-node.Val] {
			res = true
			return
		}
		m[node.Val] = true
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return res
}

//给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
func findTargetByDGAndTowPoint(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	arr := []int{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		arr = append(arr, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] < k {
			left++
		} else if arr[left]+arr[right] > k {
			right--
		} else {
			return true
		}
	}
	return false
}

//给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
func findTargetByStackAndTowPoint(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{}
	arr := []int{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, node.Val)
			root = node.Right
		}
	}

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] < k {
			left++
		} else if arr[left]+arr[right] > k {
			right--
		} else {
			return true
		}
	}
	return false
}

//检查一棵树是否是对称二叉树
//思想是 两指针操作，分别指向子树的左右
func isSymmetricByTowPointDG(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root, root)
}

func isSymmetricByIter(root *TreeNode) bool {
	l,r : =root,root
	
}