package array

import "math"

//数组中未出现的元素的集合
func FindDisappearedNumbers(nums []int) []int {
	res := []int{}
	m := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

//遍历输入数组的每个元素一次。
//我们将把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[|nums[i] |- 1] \times -1nums[∣nums[i]∣−1]×−1 。
//然后遍历数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。
//巧妙：对应位置置为负数，不影响数组对应位置的数据的判断
func FindDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		cur := int(math.Abs(float64(nums[i])))
		if nums[cur-1] > 0 {
			//nums[cur-1] = -nums[cur-1]
			nums[cur-1] *= -1
		}
	}
	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}
	return res
}