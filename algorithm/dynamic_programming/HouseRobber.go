package dynamic_programming

import (
	"math"
)

var memo []int

func rob(nums []int) int {
	memo = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	return robDG(nums, 0)
}

//自顶向下
func robDG(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}
	res := int(math.Max(
		//不抢了去下一个
		float64(robDG(nums, start+1)),
		//抢了去下下一个
		float64(nums[start]+robDG(nums, start+2))))
	memo[start] = res
	return res
}

//自底向上
func rob2(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = int(math.Max(float64(dp[i+1]), float64(nums[i]+dp[i+2])))
	}
	return dp[0]
}

//简化版
func robJ(nums []int) int {
	dpi1, dpi2 := 0, 0
	dpi := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dpi = int(math.Max(float64(dpi1), float64(nums[i]+dpi2)))
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}

//房子是一个环形
func robRange(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return int(math.Max(float64(robRangeD(nums, 0, len(nums)-2)), float64(robRangeD(nums, 1, len(nums)-1))))
}

func robRangeD(nums []int, start, end int) int {
	dpi1, dpi2 := 0, 0
	dpi := 0
	for i := end; i >= start; i-- {
		dpi = int(math.Max(float64(dpi1), float64(nums[i]+dpi2)))
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//房子是一个树形
func robTree(root *TreeNode) int {
	res := robTreeD(root)
	return int(math.Max(float64(res[0]), float64(res[1])))
}

func robTreeD(root *TreeNode) [2]int {
	//数组0位代表不抢最大值，1位代表抢的最大值
	if root == nil {
		return [2]int{0, 0}
	}
	left := robTreeD(root.Left)
	right := robTreeD(root.Right)
	//不抢
	notRob := int(math.Max(float64(left[0]), float64(left[1]))) +
		int(math.Max(float64(right[0]), float64(right[1])))
	//抢
	rob := root.Val + left[0] + right[0]
	return [2]int{notRob, rob}
}
