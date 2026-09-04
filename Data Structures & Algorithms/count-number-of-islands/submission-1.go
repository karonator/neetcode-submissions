func numIslands(grid [][]byte) int {
	var fill func(int, int)
	fill = func(x int, y int) {
		stack := [][]int{
			[]int{x, y},
		}

		for len(stack) > 0 {
			elem := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			ix, iy := elem[0], elem[1]
			grid[ix][iy] = 'X'

			if ix > 0 && grid[ix - 1][iy] == '1' {
				stack = append(stack, []int{ix - 1, iy})
			}
			if (ix < len(grid) - 1) && grid[ix + 1][iy] == '1' {
				stack = append(stack, []int{ix + 1, iy})
			}
			if iy > 0 && grid[ix][iy - 1] == '1' {
				stack = append(stack, []int{ix, iy - 1})
			}
			if (iy < len(grid[0]) - 1) && grid[ix][iy + 1] == '1' {
				stack = append(stack, []int{ix, iy + 1})
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
	fmt.Println(grid)
	return cnt
}
