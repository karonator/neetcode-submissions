func maxProfit(prices []int) int {
	INF := 1000000
	free := 0
	hold := 0
	cooldown := 0

	for i := range prices {
		if i == 0 {
			hold = -prices[i]
			cooldown = -INF
		} else {
			newHold := max(hold, free - prices[i])
			newFree := max(free, cooldown)
			newCooldown := hold + prices[i]

			free = newFree
			hold = newHold
			cooldown = newCooldown
		}
	}
	return max(cooldown, free)
}
