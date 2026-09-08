func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	data := make([][]int, len(obstacleGrid))
	for i := range obstacleGrid {
		data[i] = make([]int, len(obstacleGrid[0]))
	}

	for i := range obstacleGrid {
		for j := range obstacleGrid[0] {
			if i == 0 || j == 0 {
				if obstacleGrid[i][j] == 1 {
					data[i][j] = 0
				} else {
					if i > 0 {
						data[i][j] = data[i - 1][j]
					} else if j > 0 {
						data[i][j] = data[i][j - 1]
					} else {
						data[0][0] = 1
					}
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					data[i][j] = 0
				} else {
					data[i][j] = data[i - 1][j] + data[i][j - 1]
				}
			}
		}
	}
	return data[len(data) - 1][len(data[0]) - 1]
}
