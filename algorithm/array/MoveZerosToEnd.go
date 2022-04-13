package array

//移动0
func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			k++
		} else {
			i++
		}
	}
	for i := 0; i < k; i++ {
		nums = append(nums, 0)
	}
}