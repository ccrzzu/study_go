package dynamic_programming

import "math"

/**
 * 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
 */
func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	a, b, c := 0, 0, 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 1; i < n; i++ {
		t1, t2, t3 := res[a]*2, res[b]*3, res[c]*5
		res[i] = int(math.Min(math.Min(float64(t1), float64(t2)), float64(t3)))
		if res[i] == t1 {
			a++
		}
		if res[i] == t2 {
			b++
		}
		if res[i] == t3 {
			c++
		}
	}
	return res[n-1]
}
