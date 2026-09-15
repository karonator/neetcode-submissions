func rotate(matrix [][]int)  {
	for i := range len(matrix) / 2 {
		for j := range matrix {
			matrix[i][j], matrix[len(matrix) - i - 1][j] = matrix[len(matrix) - i - 1][j], matrix[i][j]
		}
	}

	for i := range matrix {
		for j := range matrix {
			if i > j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
}
