func pacificAtlantic(heights [][]int) [][]int {
    pacific := make([][]int, len(heights))
	atlantic := make([][]int, len(heights))
	
	for i := range pacific {
		pacific[i] = make([]int, len(heights[0]))
		atlantic[i] = make([]int, len(heights[0]))
	}

	trace := func(queue [][2]int, data [][]int) {
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for i := range queue {
			data[queue[i][0]][queue[i][1]] = 1
		}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x := cur[0] + dir[0]
				y := cur[1] + dir[1]
				if x >= 0 && x < len(data) && y >= 0 && y < len(data[0]) {
					if heights[x][y] >= heights[cur[0]][cur[1]] && data[x][y] == 0 {
						data[x][y] = 1
						queue = append(queue, [2]int{x, y})
					}
				}
			}
		}
	}

	pacificQueue := [][2]int{}
	atlanticQueue := [][2]int{}
	for i := range pacific {
		for j := range pacific[0] {
			if i == 0 || j == 0 {
				pacificQueue = append(pacificQueue, [2]int{i, j})
			}
			if i == len(pacific) - 1 || j == len(pacific[0]) - 1 {
				atlanticQueue = append(atlanticQueue, [2]int{i, j})
			}
		}
	}

	trace(pacificQueue, pacific)
	trace(atlanticQueue, atlantic)

	ans := [][]int{}

	for i := range pacific {
		for j := range pacific[0] {
			if pacific[i][j] == 1 && atlantic [i][j] == 1 {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}
