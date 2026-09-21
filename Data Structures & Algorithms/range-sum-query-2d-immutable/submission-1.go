type NumMatrix struct {
	prefixSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSums := make([][]int, len(matrix))
	for i := range prefixSums {
		prefixSums[i] = make([]int, len(matrix[0]))
	}

	for i := range matrix {
		for j := range matrix[0] {
			top := 0
			left := 0
			if i > 0 { top = prefixSums[i-1][j] }
			if j > 0 { left = prefixSums[i][j-1] }
			prefixSums[i][j] = top + left + matrix[i][j]
			if i > 0 && j > 0 {
				prefixSums[i][j] -= prefixSums[i-1][j-1]
			}
		}
	}
	return NumMatrix{
		prefixSums: prefixSums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	top := 0
	left := 0
	center := 0
	if row1 > 0 && col1 > 0 {
		center = this.prefixSums[row1 - 1][col1 - 1]
	}
	if row1 > 0 {
		top = this.prefixSums[row1 - 1][col2]
	}
	if col1 > 0 {
		left = this.prefixSums[row2][col1 - 1]
	}
	return this.prefixSums[row2][col2] - top - left + center
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
