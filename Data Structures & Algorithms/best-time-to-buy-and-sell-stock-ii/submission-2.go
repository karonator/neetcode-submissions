func maxProfit(prices []int) int {
	// монетка на руках сегодня, как я сюда пришел?
	hold := make([]int, len(prices))
	// сегодня я без монетки, как я сюда пришел?
	free := make([]int, len(prices))

	hold[0] = -prices[0]
	free[0] = 0

	for i := 1; i < len(prices); i++ {
		// либо монетка была, либо вчера не было а сегодня я ее купил
		hold[i] = max(hold[i - 1], free[i - 1] - prices[i])
		// либо монетка была вчера и я ее сегодня продал, либо вчера ее тоже не было
		free[i] = max(hold[i - 1] + prices[i], free[i - 1])
	}
	return free[len(free) - 1]
}
