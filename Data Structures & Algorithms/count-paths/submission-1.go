func uniquePaths(m int, n int) int {
    data := make([][]int, m)
	for i := range m {
		data[i] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			if i == 0 || j == 0  {
				data[i][j] = 1
			} else {
				data[i][j] = data[i - 1][j] + data[i][j - 1]
			}
		}
	}
	return data[m - 1][n - 1]
}
