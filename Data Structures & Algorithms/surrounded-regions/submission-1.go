func solve(board [][]byte) {
    fill := func(i, j int) {
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		queue := [][2]int{{i, j}}
		board[i][j] = 'A'
		for len(queue) > 0 {
			cur := queue[0] 
			queue = queue[1:]
			for _, dir := range dirs {
				x := cur[0] + dir[0]
				y := cur[1] + dir[1]
				if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
					if board[x][y] == 'O' {
						board[x][y] = 'A'
						queue = append(queue, [2]int{x, y})
					}
				}
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if i == 0 || j == 0 || i == len(board) - 1 || j == len(board[0]) - 1 {
				if board[i][j] == 'O' {
					fill(i, j)
				}
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
