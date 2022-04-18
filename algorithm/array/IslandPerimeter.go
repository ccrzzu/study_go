package array

/**
  给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。

网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
*/

func islandPerimeter(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				//上
				if i == 0 || grid[i-1][j] == 0 {
					count++
				}
				//下
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					count++
				}
				//左
				if j == 0 || grid[i][j-1] == 0 {
					count++
				}
				//右
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
					count++
				}
			}
		}
	}
	return count
}
