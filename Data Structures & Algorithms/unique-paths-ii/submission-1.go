func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	path_grid := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		path_grid[i] = make([]int, len(obstacleGrid[0]))
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 && obstacleGrid[0][0] != 1 {
				path_grid[i][j] = 1
			} else {
				if obstacleGrid[i][j] != 1 {
					top := 0
					if i > 0 {
						top = path_grid[i - 1][j]
					}
					left := 0
					if j > 0 {
						left = path_grid[i][j - 1]
					}
					path_grid[i][j] = top + left
				}
			}
		}
	}

	return path_grid[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]
}
