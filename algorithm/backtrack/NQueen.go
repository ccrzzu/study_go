package backtrack

//其实 DFS算法就是回溯算法
//(DFS,深度优先搜索算法（英語：Depth-First-Search，DFS）
//是一种用于遍历或搜索树或图的算法。)

var nQueensResult [][]string

//n皇后问题
func SolveNQueens(n int) [][]string {
	// 初始化result
	//nQueensResult = [][] string{}
	track := make([][]string, 0)
	// 初始化路径track "."，选择后改为"Q"
	for i := 0; i < n; i++ {
		tmp := make([]string, 0)
		for j := 0; j < n; j++ {
			tmp = append(tmp, ".")
		}
		track = append(track, tmp)
	}
	nQueenHelper(track, 0)
	return nQueensResult
}

func isQueenValid(track [][]string, row, col int) bool {
	//计算出来track总共有几行
	//n := len(track)
	//行不能一样
	for i := 0; i < col; i++ {
		if track[row][i] == "Q" {
			return false
		}
	}
	//列不能一样
	for i := 0; i < len(track); i++ {
		if track[i][col] == "Q" {
			return false
		}
	}
	//对角不能一样"\"
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	//对角不能一样"/"
	for i, j := row-1, col+1; i >= 0 && j < len(track); i, j = i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	return true
}

func nQueenHelper(track [][]string, row int) {
	//结束条件，row循环到最后一行
	if row == len(track) {
		tmp := make([]string, 0)
		// 将每行的选择结果改为字符串  [[".",".","Q","."],..] => ["..Q.", ...]
		for _, v := range track {
			str := ""
			for _, item := range v {
				str += item
			}
			tmp = append(tmp, str)
		}
		nQueensResult = append(nQueensResult, tmp)
		return
	}
	//计算出来现在到的列数
	n := len(track[row])
	for col := 0; col < n; col++ {
		//判断是否可以下皇后棋
		if !isQueenValid(track, row, col) {
			continue
		}
		//选择下皇后
		track[row][col] = "Q"
		//继续下次决策选择
		nQueenHelper(track, row+1)
		//撤回选择
		track[row][col] = "."
	}
}
