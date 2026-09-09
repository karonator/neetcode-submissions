func coinChange(coins []int, amount int) int {
	data := make([]int, amount + 1)
	data[0] = 0
	for i := 1; i <= amount; i++ {
		data[i] = amount + 1
	}
 
	for i := 1; i <= amount; i++ {
		for j := range coins {
			if i - coins[j] >= 0 {
				data[i] = min(data[i], data[i - coins[j]] + 1)
			}
		}
	}

	if data[amount] == amount + 1 {
		return -1
	}
	return data[amount]
}
