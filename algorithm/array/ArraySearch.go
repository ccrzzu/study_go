package array

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
