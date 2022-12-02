package tree

/**
DFS : depth first search 深度优先遍历

四种主要的遍历思想为：

前序（先根）遍历：根结点 ---> 左子树 ---> 右子树

中序（中根）遍历：左子树 ---> 根结点 ---> 右子树

后序（后根）遍历：左子树 ---> 右子树 ---> 根结点

层次遍历：只需按层次遍历即可，属于BFS（breadth first search）了
*/

//二叉树的前序遍历 递归
func preOrderTraversalByDG(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		leftRes := preOrderTraversalByDG(root.Left)
		res = append(res, leftRes...)
		rightRes := preOrderTraversalByDG(root.Right)
		res = append(res, rightRes...)
	}
	return res
}

//二叉树的前序遍历 非递归，用栈模拟递归过程
func preOrderTraversalByStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}
	stack = append(stack, root)
	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, lastNode.Val)
		if lastNode.Right != nil {
			stack = append(stack, lastNode.Right)
		}
		if lastNode.Left != nil {
			stack = append(stack, lastNode.Left)
		}
	}
	return res
}

//二叉树的中序遍历 递归
func inOrderTraversalByDG(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

//二叉树的中序遍历 非递归，用栈模拟递归过程
func inOrderTraversalByStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root.Left != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			root = node.Right
		}
	}
	return res
}

//二叉树的后序遍历 递归
func postOrderTraversalByDG(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		leftRes := preOrderTraversalByDG(root.Left)
		res = append(res, leftRes...)
		rightRes := preOrderTraversalByDG(root.Right)
		res = append(res, rightRes...)
		res = append(res, root.Val)
	}
	return res
}

//二叉树的后序遍历 非递归，用栈模拟递归过程,
//前序是 根左右 稍微改成 根右左，然后反转就是后序 左右根
func postOrderTraversalByStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}
	stack = append(stack, root)
	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, lastNode.Val)
		if lastNode.Left != nil {
			stack = append(stack, lastNode.Left)
		}
		if lastNode.Right != nil {
			stack = append(stack, lastNode.Right)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
