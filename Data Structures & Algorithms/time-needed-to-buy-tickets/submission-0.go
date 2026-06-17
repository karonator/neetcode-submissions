func timeRequiredToBuy(tickets []int, k int) int {
    result := 0
	runner := 0
	for tickets[k] != 0 {
		if tickets[runner] > 0 {
			tickets[runner]--
			result++
		}
		runner++
		if runner == len(tickets) {
			runner = 0
		}
	}
	return result
}