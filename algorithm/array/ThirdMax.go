package array

import (
	"fmt"
	"math"
)

/**
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是 O(n)。
*/
func ThirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, item := range nums {
		if item > first {
			third = second
			second = first
			first = item
		} else if item < first && item > second {
			third = second
			second = item
		} else if item < second && item > third {
			third = item
		}
		fmt.Println(first, second, third)
	}
	if third == math.MinInt64 {
		return first
	}
	return third
}
