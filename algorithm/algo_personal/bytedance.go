package main

/**
leetcode 207
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，
表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
变种题：
给你一个二维数组，判断是否有环
[[a,b],[b,c],....[a,d]]

解题思路：
见graph模块
**/
// AOV 网的拓扑排序
func canFinish(n int, pre [][]int) bool {
	in := make([]int, n)
	frees := make([][]int, n)
	next := make([]int, 0, n)
	for _, v := range pre {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	return len(next) == n
}

/**
一个有序的集合，判断一个元素是否重复量超过总数的一半
**/
func checkBeyondHalf(arr []int, target int) bool {

	return true
}

/**
给定一个无序的整数数组，其中数据可能重复，要求将其中的奇数排序并放置到数组头部，
偶数则维持当前顺序放置到数组后部，要求不能创建新数组
例：[1,2,3,1,6,3,4,9]
结果：[1,1,3,3,9,2,6,4]
思想：
1、先把偶数都放到最后
2、把前面的数字再来个快速排序即可
**/
//TODO 待梳理完善
func sort(nums []int) []int {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		numsi := nums[i]
		numsj := nums[i-1]
		//前奇后偶，交换
		if numsi%2 != 0 && numsj%2 == 0 {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		} else if numsi%2 != 0 && numsj%2 != 0 {
			//前奇后奇且前大后小，交换
			if numsj > numsi {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return nums
}
