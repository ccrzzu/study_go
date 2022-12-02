package tree

/**
199
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
解题思路：
这一题是按层序遍历的变种题。按照层序把每层的元素都遍历出来，然后依次取每一层的最右边的元素即可。用一个队列即可实现。
**/
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		res = append(res, queue[qLen-1].Val)
		for i := 0; i < qLen; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[qLen:]
	}
	return res
}

/**
左视图
**/
func leftSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		res = append(res, queue[0].Val)
		for i := 0; i < qLen; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[qLen:]
	}
	return res
}
