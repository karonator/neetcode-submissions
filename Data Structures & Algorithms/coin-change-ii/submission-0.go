func change(amount int, coins []int) int {
    data := make([][]int, len(coins) + 1)
	for i := range data {
		data[i] = make([]int, amount + 1)
		data[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			data[i][j] = data[i - 1][j]
			if j - coins[i - 1] >= 0 {
				data[i][j] += data[i][j - coins[i - 1]]
			}
		}
	}

	return data[len(coins)][amount]
}
