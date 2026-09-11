func islandsAndTreasure(grid [][]int) {
	INF := math.MaxInt32

	queue := [][2]int{}
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}

	steps := 1
	for len(queue) > 0 {
		stash := [][2]int{}
		for _, coord := range queue {
			for _, dir := range dirs {
				x := coord[0] + dir[0]
				y := coord[1] + dir[1]
				if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) {
					if grid[x][y] == INF {
						grid[x][y] = steps
						stash = append(stash, [2]int{x, y})
					}
				}
			}
		}
		steps++
		queue = stash
	}
}
