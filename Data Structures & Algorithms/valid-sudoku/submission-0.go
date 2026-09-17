func isValidSudoku(board [][]byte) bool {
	rows := make([]map[int]struct{}, 9)
	cols := make([]map[int]struct{}, 9)
	quads := make([]map[int]struct{}, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]struct{})
		cols[i] = make(map[int]struct{})
		quads[i] = make(map[int]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, found := rows[i][int(board[i][j] - '0')]; found {
					return false
				}
				rows[i][int(board[i][j] - '0')] = struct{}{}

				if _, found := cols[j][int(board[i][j] - '0')]; found {
					return false
				}
				cols[j][int(board[i][j] - '0')] = struct{}{}

				quadCoord := i / 3 + 3 * (j / 3)
				if _, found := quads[quadCoord][int(board[i][j] - '0')]; found {
					return false
				}				
				quads[quadCoord][int(board[i][j] - '0')] = struct{}{}
			}
		}
	}
	return true
}
