func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	ans := make([][]int, len(grid))
	for i := range(grid) {
		ans[i] = make([]int, len(grid[0]))
	}
	ans[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				ans[i][j] = grid[i][j] + ans[i][j - 1]
				continue
			}
			if j == 0 {
				ans[i][j] = grid[i][j] + ans[i - 1][j]
				continue
			}
			ans[i][j] = grid[i][j] + min(ans[i - 1][j], ans[i][j - 1])
		}
	}

	return ans[len(grid) - 1][len(grid[0]) - 1]
}
