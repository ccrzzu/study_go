package tree


//判断一棵树是否为完全二叉树
//完全二叉树的定义就是每一层从左往右都是连续的，不能有空的跳过去
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	lastIsNil := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			lastIsNil = true
		} else {
			if lastIsNil {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}