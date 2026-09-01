func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	lo, hi := 0, rows * cols - 1
	for lo < hi {
		mid := lo + (hi - lo) / 2
		row := mid / cols
		col := mid - row * cols

		if target <= matrix[row][col]  {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	row := lo / cols
	col := lo - row * cols

	if matrix[row][col] == target {
		return true
	}

	return false
}
