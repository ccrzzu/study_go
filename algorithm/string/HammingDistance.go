package string

import "math/bits"

//汉明距离
//两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//思路：位运算里，异或指的是：两个位相同为0,不同为1
//汉明距离广泛应用于多个领域。在编码理论中用于错误检测，在信息论中量化字符串之间的差异。
func hammingDistance(x int, y int) int {
	var sum int
	i := x ^ y
	for i != 0 {
		if i&1 == 1 {
			sum++
		}
		i = i >> 1
	}
	return sum
	return bits.OnesCount(uint(x^y))
}
