package tree


/**
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
**/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := math.MinInt64
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	root := &TreeNode{
		Val: maxVal,
	}
	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}