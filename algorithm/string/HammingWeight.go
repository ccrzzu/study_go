package string

import "math/bits"

// 解法一
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

// 解法二
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// 解法三
func hammingWeight3(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
