package array

import (
	"fmt"
	"strconv"
)

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

// 注意这种解法对于数组长度过长后就不可以了，这样写会越界，小于int64表示的数的长度还可以
func PlusOne2(digits []int) []int {
	var numsStr string
	for _, item := range digits {
		numsStr += strconv.Itoa(item)
	}
	fmt.Println(numsStr)
	res := []int{}
	numInt, _ := strconv.Atoi(numsStr)
	for _, v := range strconv.Itoa(numInt + 1) {
		fmt.Println(v - '0')
		res = append(res, int(v-'0'))
	}
	fmt.Println(res)
	return res
}

/**
*整数的 数组形式  num 是按照从左到右的顺序表示其数字的数组。

例如，对于 num = 1321 ，数组形式是 [1,3,2,1] 。
给定 num ，整数的 数组形式 ，和整数 k ，返回 整数 num + k 的 数组形式 。
*/
func AddToArrayForm(num []int, k int) []int {
	res := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			sum = sum % 10
			k++
		}
		res = append(res, sum)
	}
	for ; k > 0; k /= 10 {
		res = append(res, k%10)
	}
	ReverseArray(res)
	fmt.Println(res)
	return res
}
