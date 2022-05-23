package array

import (
	"math"
	"sort"
)

/**
 * 总：解决区间问题的一般思路是先排序，然后观察规律。
 * 一般会用到双指针。
 */

/**
1288
给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
在完成所有删除操作后，请你返回列表中剩余区间的数目
*/
func RemoveCoveredIntervals(intervals [][]int) int {
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

/**
 56
 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func MergeInterval(intervals [][]int) [][]int {
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

/**
986
给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，
其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。
每个区间列表都是成对 不相交 的，并且 已经排序 。

返回这 两个区间列表的交集 。
形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
*/
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
