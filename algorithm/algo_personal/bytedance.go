package main

/**
有多少种零食组合
**/







/**
给定一个无序的整数数组，其中数据可能重复，要求将其中的奇数排序并放置到数组头部，
偶数则维持当前顺序放置到数组后部，要求不能创建新数组
例：[1,2,3,1,6,3,4,9]
结果：[1,1,3,3,9,2,6,4]
思想：
1、先把偶数都放到最后
2、把前面的数字再来个快速排序
**/
//TODO
func sort(nums []int) []int{
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		numsi := nums[i]
		numsj := nums[i-1]
		//前奇后偶，交换
		if numsi%2 != 0 && numsj%2 == 0 {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}else if numsi%2 != 0 && numsj%2 != 0{
			//前奇后奇且前大后小，交换
			if numsj > numsi{
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return nums
}