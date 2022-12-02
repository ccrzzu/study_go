package array

func longestLine(M [][]byte) int {
	if len(M) == 0 || len(M[0]) == 0 {
		return 0
	}
	res, m, n := 0, len(M), len(M[0])
	row := make([]int, m)
	col := make([]int, n)
	d := make([]int, m+n)
	ad := make([]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 'F' {
				row[i]++
				col[j]++
				d[i+j]++
				ad[j-i+m]++
				res = max(res, max(row[i], col[j]))
				res = max(res, max(d[i+j], ad[j-i+m]))
			} else {
				row[i] = 0
				col[j] = 0
				d[i+j] = 0
				ad[j-i+m] = 0
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
