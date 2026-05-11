type NumMatrix struct {
    sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    sums := make([][]int, len(matrix))
    for i := range len(matrix) {
        sums[i] = make([]int, len(matrix[i]))
        for j := range len(matrix[i]) {
            prefixTop := 0
            if i > 0 {
                prefixTop = sums[i - 1][j]
            }
            prefixLeft := 0
            if j > 0 {
                prefixLeft = sums[i][j - 1]
            }
            prefixCenter := 0
            if i > 0 && j > 0 {
                prefixCenter = sums[i - 1][j - 1]
            }
            sums[i][j] = prefixTop + prefixLeft + matrix[i][j] - prefixCenter
        }
    }
    return NumMatrix {
        sums: sums,
    }
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    divTop := 0
    if row1 > 0 {
        divTop = this.sums[row1 - 1][col2]
    }

    divLeft := 0
    if col1 > 0 {
        divLeft = this.sums[row2][col1 - 1]
    }

    divCenter := 0
    if row1 > 0 && col1 > 0 {
        divCenter = this.sums[row1 - 1][col1 - 1]
    }

    return this.sums[row2][col2] - divTop - divLeft + divCenter
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)


