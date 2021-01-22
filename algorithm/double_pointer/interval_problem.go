package double_pointer

import (
	"math"
	"sort"
)

/**
 * 总：解决区间问题的一般思路是先排序，然后观察规律。
 */

//移除覆盖区间，得到实际区间个数
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	//记录起点和终点
	left, right := intervals[0][0], intervals[0][1]
	var res int
	for i := 1; i < len(intervals); i++ {
		item := intervals[i]
		//情况1，找到覆盖区间了
		if item[0] >= left && item[1] <= right {
			res++
		}
		//情况2，区间交叉
		if item[0] >= left && item[1] > right {
			right = item[1]
		}
		//情况3，区间完全不相交
		if item[0] > right {
			left, right = item[0], item[1]
		}
	}
	return len(intervals) - res
}

//合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		curr := intervals[i]
		if curr[0] <= last[1] {
			last[1] = int(math.Max(float64(curr[1]), float64(last[1])))
		} else {
			res = append(res, curr)
		}
	}
	return res
}

//两个有序区间列表的交集
func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	res := [][]int{}
	for i < len(A) && j < len(B) {
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]
		if a1 <= b2 && a2 >= b1 {
			res = append(res, []int{int(math.Max(float64(a1), float64(b1))), int(math.Min(float64(a2), float64(b2)))})
		}
		if a2 < b2 {
			i++
		} else {
			j++
		}
	}
	return res
}
