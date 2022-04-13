package array

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
func MoreThanHalfNum(numbers []int) int {
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