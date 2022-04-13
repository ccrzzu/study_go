package array

//加一操作
/**
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
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