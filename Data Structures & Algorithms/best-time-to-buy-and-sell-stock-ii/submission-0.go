func maxProfit(prices []int) int {
	min := -1
	profit := 0
	potential_profit := 0

	for _, price := range prices {
		if min == -1 || min > price {
			min = price
			profit += potential_profit
			potential_profit = 0
			continue
		}
		current_potential_profit := price - min
		if current_potential_profit > potential_profit {
			potential_profit = current_potential_profit
			continue
		} else {
			profit += potential_profit
			potential_profit = 0
			min = price
		}
	}
	profit += potential_profit

	return profit
}
