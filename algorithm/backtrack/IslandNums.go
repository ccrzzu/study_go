package backtrack

/**
200
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

例子
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
*/

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 递归return 越界或遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}
		// 已访问过 说明到接壤处了 也返回
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i-1, j) //上
		dfs(i+1, j) //下
		dfs(i, j-1) //左
		dfs(i, j+1) //右
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
