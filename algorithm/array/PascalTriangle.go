package array

// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。在杨辉三角中，每个数是它左上方和右上方的数的和。
func PascalTriangleGenerate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		jNum := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				jNum = append(jNum, 1)
			} else {
				jNum = append(jNum, res[i-1][j-1]+res[i-1][j])
			}
		}
		res = append(res, jNum)
	}
	return res
}

// 给定一个非负整数 rowIndex，只返回第rowIndex行的一维数组
// 题目要求我们只能使用 O(k) 的空间。那么需要找到两两项直接的递推关系。
func PascalTriangleGetRow(rowIndex int) []int{
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
