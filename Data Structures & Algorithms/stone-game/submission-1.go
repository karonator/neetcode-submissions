func stoneGame(piles []int) bool {
	n := len(piles)
	data := make([][]int, n)
	for i := range data {
		data[i] = make([]int, n)
	}
	for i := range data {
		data[i][i] = piles[i]
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 1; r < n; r++ {
			data[l][r] = max(piles[l] - data[l + 1][r], piles[r] - data[l][r - 1])
		}
	}
	return data[0][n-1] > 0
}
