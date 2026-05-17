func uniquePaths(m int, n int) int {
    grid := make([][]int, m)
	for i := range(m) {
		grid[i] = make([]int, n)
	}
	grid[0][0] = 1

	for i := range(m) {
		for j := range(n) {
			if i == 0 && j == 0 {
				continue
			}
			top := 0
			if i > 0 {
				top = grid[i - 1][j]
			}
			left := 0
			if j > 0 {
				left = grid[i][j - 1]
			}
			grid[i][j] = top + left
		}
	}

	return grid[m - 1][n - 1]
}
