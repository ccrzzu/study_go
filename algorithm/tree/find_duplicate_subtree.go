package tree

import (
	"fmt"
	"strconv"
)

//寻找重复的子树
var subTreesMap map[string]int
var dRes []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTreesMap = map[string]int{}
	dRes = make([]*TreeNode, 0)
	findDuplicateSubtreesDG(root)
	return dRes
}

//寻找重复的子树 递归解法
func findDuplicateSubtreesDG(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := findDuplicateSubtreesDG(root.Left)
	right := findDuplicateSubtreesDG(root.Right)
	subTreeStr := left + "," + right + "," + strconv.Itoa(root.Val)
	subTreesMap[subTreeStr]++
	if subTreesMap[subTreeStr] == 2 {
		dRes = append(dRes, root)
	}
	fmt.Println(subTreesMap, dRes)
	return subTreeStr
}