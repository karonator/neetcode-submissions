func islandsAndTreasure(grid [][]int) {
	INF := 2147483647

    queue := [][]int{}
	for i := range(grid) {
		for j := range(grid[0]) {
			if grid[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	fmt.Println(queue)
	dist := 1
	for len(queue) > 0 {
		tmp := [][]int{}
		for _, v := range(queue) {
			if v[0] > 0 && grid[v[0] - 1][v[1]] == INF {
				grid[v[0] - 1][v[1]] = dist
				tmp = append(tmp, []int{v[0] - 1,v[1]})
			}
			if v[1] > 0 && grid[v[0]][v[1] - 1] == INF {
				grid[v[0]][v[1] - 1] = dist
				tmp = append(tmp, []int{v[0],v[1] - 1})
			}
			if v[0] + 1 < len(grid) && grid[v[0] + 1][v[1]] == INF {
				grid[v[0] + 1][v[1]] = dist
				tmp = append(tmp, []int{v[0] + 1,v[1]})
			}
			if v[1] + 1 < len(grid[0]) && grid[v[0]][v[1] + 1] == INF {
				grid[v[0]][v[1] + 1] = dist
				tmp = append(tmp, []int{v[0],v[1] + 1})
			}
		}
		queue = tmp
		dist++
	}
}

// L W X L
// L L L W
// L W L W
// X W L L