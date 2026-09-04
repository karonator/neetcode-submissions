func numIslands(grid [][]byte) int {
	fill := func(x int, y int) {
		stack := [][2]int{
			[2]int{x, y},
		}

		for len(stack) > 0 {
			elem := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			ix, iy := elem[0], elem[1]
			grid[ix][iy] = 'X'
			dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

			for _, dir := range(dirs) {
				jx := ix + dir[0]
				jy := iy + dir[1]
				if jx >= 0 && jy >= 0 && jx < len(grid) && jy < len(grid[0]) {
					if grid[jx][jy] == '1' {
						stack = append(stack, [2]int{jx, jy})
					}
				}
			}

		}
	}

	cnt := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				fill(i, j)
				cnt ++
			}
		}
	}
	return cnt
}
