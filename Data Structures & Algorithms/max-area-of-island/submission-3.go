func maxAreaOfIsland(grid [][]int) int {
    fill := func(x int, y int) int {
		stack := [][]int{{x, y}}
		area := 0
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for len(stack) > 0 {
			point := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			// важно, так как иначе одна и та же клетка посчитается дважы если на нее указали
			// две соседние клетки
			if grid[point[0]][point[1]] == 1 {
				area++
				grid[point[0]][point[1]] = 2
				for _, dir := range dirs {
					x := point[0] + dir[0]
					y := point[1] + dir[1]
					if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1 {
						stack = append(stack, []int{x, y})
					}
				}
			}
		}
		return area
	}

	maxArea := 0
	for	i := range(grid) {
		for j := range(grid[0]) {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, fill(i, j))
			}
		}
	}

	return maxArea
}
