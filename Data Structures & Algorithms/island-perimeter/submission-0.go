func islandPerimeter(grid [][]int) int {
	result := 0
	for i := range(grid) {
		for j := range(grid[0]) {
			if grid[i][j] == 1 {
				if i == 0 {
					result ++
				} else {
					if grid[i - 1][j] == 0 {
						result ++
					}
				}
				if j == 0 {
					result ++
				} else {
					if grid[i][j - 1] == 0 {
						result ++
					}
				}
				if i == len(grid) - 1 {
					result ++
				} else {
					if grid[i + 1][j] == 0 {
						result ++
					}
				}
				if j == len(grid[0]) - 1 {
					result ++
				} else {
					if grid[i][j + 1] == 0 {
						result ++
					}
				}
			}
		}
	}
	return result
}
