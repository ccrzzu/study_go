package double_pointer

//两数相加等于目标数
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target{
			return []int{left+1,right+1}
		}else if sum < target{
			left++
		}else if sum > target{
			right--
		}
	}
	return []int{-1,-1}
}

//反转字符串
func reverseString(s []byte)  {
	left,right := 0, len(s)-1
	for left < right{
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}