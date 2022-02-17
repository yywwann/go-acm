package knight_probability_in_chessboard

var dirs = []struct{ i, j int }{
	{-2, -1},
	{-2, 1},
	{2, -1},
	{2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
}

func knightProbability(n, k, row, column int) float64 {
	dp := make([][][]float64, k+1) // dp[step][i][j] 表示从i,j开始骑士走step步仍在棋盘内的概率
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
