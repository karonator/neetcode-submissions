func orangesRotting(grid [][]int) int {
	queue := make([][2]int, 0)

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	steps := 0
	for len(queue) > 0 {
		tmp := make([][2]int, 0)
		for i := range queue {
			for _, dir := range dirs {
				x := queue[i][0] + dir[0]
				y := queue[i][1] + dir[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					tmp = append(tmp, [2]int{x, y})
					grid[x][y] = 2
				}
			}
		}
		queue = tmp
		if len(queue) > 0 {
			steps++
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return steps
}
