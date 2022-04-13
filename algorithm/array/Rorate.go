package array

// 轮转数组
// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Rotate1(nums []int, k int) {
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[0:k]...))
}

func Rotate2(nums []int, k int) {
	//使用channel 队列FIFO方式，空间复杂度O(n),时间复杂度O(n+m)
	num := make(chan int, len(nums))
	m := k % len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		num <- nums[i]
	}
	for i := 0; i < m; i++ {
		numVal, _ := <-num
		num <- numVal
	}
	close(num)
	i := len(num) - 1
	for k := range num {
		nums[i] = k
		i--
	}
}


func Rotate3(arr []int, k int) {
	ReverseArray(arr)
	ReverseArray(arr[:k%len(arr)])
	ReverseArray(arr[k%len(arr):])
}