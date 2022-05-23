package array


//验证是否为山峰数组
func ValidMountainArray(A []int) bool {
	i := 0
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for i+1 < len(A) && A[i] > A[i+1] {
		i++
	}

	return i == len(A)-1
}