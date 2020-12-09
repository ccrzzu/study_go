package array

//旋转数组
func rotateArr(arr []int, k int) {
	reverseArray(arr)
	reverseArray(arr[:k%len(arr)])
	reverseArray(arr[k%len(arr):])
}

func rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[0:k]...))
}

func rotate2(nums []int, k int)  {
	//使用channel 队列FIFO方式，空间复杂度O(n),时间复杂度O(n+m)
	num := make(chan int, len(nums))
	m := k % len(nums)
	for i:=len(nums)-1;i>=0;i-- {
		num <- nums[i]
	}
	for i:=0;i<m;i++ {
		numVal,_ := <- num
		num <- numVal
	}
	close(num)
	i := len(num)-1
	for k := range num {
		nums[i]=k
		i--
	}
}

//在一个旋转过的数组中查找target值
func ArrSearch( A []int ,  target int ) int {
	// write code here
	left,right := 0,len(A)-1
	for left <= right{
		mid := left + (right - left) / 2
		if A[mid] == target{
			return mid
		}
		if A[mid] < A[left]{
			//检查右边有序集
			if A[mid] < target && A[right] >= target{
				left = mid + 1
			}else {
				right = mid - 1
			}
		}else {
			//检查左边有序集
			if A[mid] > target && A[left] <= target{
				right = mid - 1
			}else {
				left = mid + 1
			}
		}
	}
	return -1
}



//反转数组
func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
