package tree

//将满二叉树水平层相邻节点指向下一个
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 将传入的两个节点连接
	node1.Next = node2
	// 连接相同父节点的两个子节点
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)

	//连接不同父节点的两个子节点
	connectTwoNode(node1.Right, node2.Left)
}
