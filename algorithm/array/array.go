package array

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

//旋转数组
func rotateArr(arr []int, k int) {
	ReverseArray(arr)
	ReverseArray(arr[:k%len(arr)])
	ReverseArray(arr[k%len(arr):])
}

func rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[0:k]...))
}

func rotate2(nums []int, k int) {
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

//在一个旋转过的数组中查找target值
func ArrSearch(A []int, target int) int {
	// write code here
	left, right := 0, len(A)-1
	for left <= right {
		mid := left + (right-left)/2
		if A[mid] == target {
			return mid
		}
		if A[mid] < A[left] {
			//检查右边有序集
			if A[mid] < target && A[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			//检查左边有序集
			if A[mid] > target && A[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

//反转数组
func ReverseArray(arr []int) []int{
	//第一种思路：
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	/*//第二种思路：
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}*/
	return arr
}

//加一操作
func plusOne(digits []int) []int {
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- {
		//最后一位肯定要加1
		if i == len(digits)-1 {
			digits[i]++
		} else {
			//判断非最后一位要不要+1
			digits[i] += addon
			addon = 0
		}
		//非最后一位，如果等于10，要+1
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	if addon == 1 {
		result = []int{1}
		result = append(result, digits...)
	} else {
		result = digits
	}
	return result
}

//两个数组的交集，结果中是需要包含重复次数的
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item]++
	}
	for _, item := range nums2 {
		if mp1[item] > 0 {
			mp1[item]--
			res = append(res, item)
		}
	}
	return res
}

//让两个数组排序后，然后用双指针策略
func intersectBySort(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return res
}

//让两个数组排序后，然后用双指针策略
func intersectBySortNoNewMem(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums1[k] = nums2[j]
			i++
			j++
			k++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return nums1[:k]
}

//两个数组的交集，结果中重复次数只返回一次
func intersectUnique(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item] = 0
	}
	for _, item := range nums2 {
		if _, ok := mp1[item]; ok {
			res = append(res, item)
			delete(mp1, item)
		}
	}
	return res
}

//原地删除
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//删除有序数组的重复元素
func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

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

//字符串数组的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, item := range strs {
		for strings.Index(item, prefix) != 0 {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

//合并两个排序的数组，合并B到A，形成一个新的有序数组
func Merge(A []int, m int, B []int, n int) {
	var res []int
	var i, j int
	for i < m && j < n {
		if A[i] <= B[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
		fmt.Println(res)
	}
	if i < m {
		res = append(res, A[i:m]...)
	}
	if j < n {
		res = append(res, B[j:n]...)
	}
	for i, item := range res {
		A[i] = item
	}
}

//归并排序逻辑
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, tail := m-1, n-1, m+n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[tail] = nums1[i1]
			i1--
		} else {
			nums1[tail] = nums2[i2]
			i2--
		}
		tail--
	}
	for tail >= 0 && i1 >= 0 {
		nums1[tail] = nums1[i1]
		i1--
		tail--
	}
	for tail >= 0 && i2 >= 0 {
		nums1[tail] = nums2[i2]
		i2--
		tail--
	}
}

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

//数组的最大子序列的和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//思路：动态规划的思想，对于i位置，其子序列和最大值为加上i-1的位置的子序列和最大值 和 其本身的比较
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
func MoreThanHalfNum_Solution(numbers []int) int {
	m := map[int]int{}
	n := len(numbers) / 2
	for i := 0; i < len(numbers); i++ {
		m[numbers[i]]++
		if m[numbers[i]] > n {
			return numbers[i]
		}
	}
	return 0
}
