package dynamic_programming

import (
	"math"
)

var memo1 map[[2]int]int
func SuperEggDrop(K int, N int) int {
	memo1 = make(map[[2]int]int)
	return dp(K, N)
}

func dp(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if _, ok := memo1[[2]int{K, N}]; ok {
		return memo1[[2]int{K, N}]
	}
	low, high := 1, N
	var res = math.MaxInt32
	for low <= high {
		mid := (low + high) / 2
		broken := dp(K-1, mid-1)    //碎
		noBroken := dp(K, N-mid) //没碎
		if broken > noBroken {
			high = mid - 1
			res = int(math.Min(float64(res), float64(broken+1)))
		} else {
			low = mid + 1
			res = int(math.Min(float64(res), float64(noBroken+1)))
		}
	}
	memo1[[2]int{K, N}] = res
	return res
}
