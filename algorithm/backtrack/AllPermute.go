package backtrack

import "fmt"

//其实 DFS算法就是回溯算法
//(DFS,深度优先搜索算法（英語：Depth-First-Search，DFS）
//是一种用于遍历或搜索树或图的算法。)

/**
46
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

// dfs回溯解法
func Permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	result := [][]int{}
	
	// nums: 原始列表 pathNums: 路径上的数字 used: 是否访问过
	var dfs func(nums, pathNums []int, used []bool)
	dfs = func(nums, pathNums []int, used []bool) {
		// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
		if len(nums) == len(pathNums) {
			tmp := make([]int, len(nums))
			// 切片底层公用数据，所以要copy
			copy(tmp, pathNums)
			// 把本次结果追加到最终结果上
			result = append(result, tmp)
			return
		}

		// 开始遍历原始数组的每个数字
		for i := 0; i < len(nums); i++ {
			// 检查是否访问过
			if !used[i] {
				// 没有访问过就选择它，然后标记成已访问过的
				used[i] = true
				// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
				pathNums = append(pathNums, nums[i])
				dfs(nums, pathNums, used)
				// 撤销刚才的选择，也就是恢复操作
				pathNums = pathNums[:len(pathNums)-1]
				// 标记成未使用
				used[i] = false
			}
		}
	}

	dfs(nums, pathNums, used)
	return result
}

var res [][]int

//全排列 递归解法
func Permutations(nums []int) [][]int {
	res = make([][]int, 0)
	track := make([]int, 0)
	permutationsDG(nums, track)
	return res
}

// 递归部分
func permutationsDG(nums, track []int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, track)
		res = append(res, tmp)
		fmt.Println("res add:", track, "result:", res)
		return
	}

	for i := 0; i < len(nums); i++ {
		if InSlice(nums[i], track) {
			//fmt.Println("track add continue:",i,track)
			continue
		}
		track = append(track, nums[i])
		fmt.Println("track add success:", i, track)
		permutationsDG(nums, track)
		track = track[:len(track)-1]
		fmt.Println("track remove:", i, track)
	}
}

func InSlice(v int, sl []int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}
