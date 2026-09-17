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
				num := int(board[i][j] - '0')
				if _, found := rows[i][num]; found {
					return false
				}
				rows[i][num] = struct{}{}

				if _, found := cols[j][num]; found {
					return false
				}
				cols[j][num] = struct{}{}

				quadCoord := i / 3 + 3 * (j / 3)
				if _, found := quads[quadCoord][num]; found {
					return false
				}				
				quads[quadCoord][num] = struct{}{}
			}
		}
	}
	return true
}
