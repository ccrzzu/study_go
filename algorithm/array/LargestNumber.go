package array

import (
	"sort"
	"strconv"
)

/**
179
给出一个数组，要求排列这些数组里的元素，使得最终排列出来的数字是最大的。
*/
func LargestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

/**
给定一个数 n，如 23121；给定一组数字 nums 如 {2,4,9}，
求由 nums 中元素组成的、小于 n 的最大数，如小于 23121 的最大数为 22999
思路：回溯法
从第一位开始进行查询,
如果能够在数组中找到等于该位的元素或者能取到小于该位元素的最大值就直接取
如果取不到,那么就少一位,以后直接每回都取数组中最大的元素即可
*/
func LargestNumberLessN(nums []int, n int) int {
	sort.Ints(nums)
	nStr := strconv.Itoa(n)
	var ans int

	//pass表示当前位要不要直接跳过
	var dfs func(index int, pass bool, temp int, arr []int) bool
	dfs = func(index int, pass bool, temp int, arr []int) bool {
		if index == len(nStr) {
			ans = temp
			return true
		}

		//pass为true的情况,只有最高位没有数，或者上一位数小于对应n中的值
		if pass {
			return dfs(index+1, true, temp*10+arr[len(arr)-1], arr)
		} else {
			//定义当前深度的n的数
			val := int(nStr[index] - '0')

			//找到与这一位相比 取相等的或者小于中最大的一位
			for i := len(arr) - 1; i >= 0; i-- {
				//先判断等于这一位的
				if arr[i] == val {
					if dfs(index+1, false, temp*10+arr[i], arr) {
						return true
					}
				} else if arr[i] < val {
					//要不就取小于这一位的最大值
					if dfs(index+1, true, temp*10+arr[i], arr) {
						return true
					}
				}
			}

			//如果该位置都放不了且index!=0
			if index != 0 {
				return false
			}

			//如果index==0 上面for循环中这些位置都不满足,应该跳过这一位了
			return dfs(index+1, true, temp, arr)
		}
	}

	dfs(0, false, 0, nums)
	return ans
}
