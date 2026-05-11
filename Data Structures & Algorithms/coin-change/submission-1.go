func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
    ans := make([]int, amount + 1)
	for i := range(ans) {
		ans[i] = math.MaxInt32
	}
	ans[0] = 0

	for i := range(ans) {
		for _, coin := range(coins) {
			if i - coin >= 0 {
				ans[i] = min(ans[i], ans[i - coin] + 1)
			}
		}
	}
	if ans[amount] == math.MaxInt32 {
		return -1
	}
	return ans[amount]
}
