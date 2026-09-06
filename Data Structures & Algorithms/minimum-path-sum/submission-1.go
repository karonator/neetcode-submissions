func minPathSum(grid [][]int) int {
	ans := make([][]int, len(grid))
	for i := range(grid) {
		ans[i] = make([]int, len(grid[0]))
	}

	for i := range(grid) {
		for j := range(grid[0]) {
			if i == 0 && j == 0 {
				ans[i][j] = grid[i][j]
			} else if i == 0 {
				ans[i][j] = ans[i][j - 1] + grid[i][j]
			} else if j == 0 {
				ans[i][j] +=  ans[i - 1][j] + grid[i][j]
			} else {
				ans[i][j] = min(ans[i - 1][j], ans[i][j - 1]) + grid[i][j]
			}
		}
	}
	return ans[len(grid) - 1][len(grid[0]) - 1]
}
