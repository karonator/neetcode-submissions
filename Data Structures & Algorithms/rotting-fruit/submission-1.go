func orangesRotting(grid [][]int) int {
    queue := [][]int{}
	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	minutes := 0
	for len(queue) > 0 {
		tmp := [][]int{}
		
		for _, v := range(queue) {
			if v[0] > 0 && grid[v[0] - 1][v[1]] == 1 {
				grid[v[0] - 1][v[1]] = 2
				tmp = append(tmp, []int{v[0] - 1, v[1]})
			}
			if v[0] < len(grid) - 1 && grid[v[0] + 1][v[1]] == 1 {
				grid[v[0] + 1][v[1]] = 2
				tmp = append(tmp, []int{v[0] + 1, v[1]})
			}
			if v[1] > 0 && grid[v[0]][v[1] - 1] == 1 {
				grid[v[0]][v[1] - 1] = 2
				tmp = append(tmp, []int{v[0], v[1] - 1})
			}
			if v[1] < len(grid[0]) - 1 && grid[v[0]][v[1] + 1] == 1 {
				grid[v[0]][v[1] + 1] = 2
				tmp = append(tmp, []int{v[0], v[1] + 1})
			}
		}
		queue = tmp
		if len(tmp) > 0 {
			minutes++
		}
	}

	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
