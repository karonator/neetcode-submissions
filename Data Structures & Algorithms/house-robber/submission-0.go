func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	money := -1

	pre_prev := 0
	prev := 0

	for _, num := range(nums) {
		money = max(money, num + pre_prev)
		pre_prev = prev
		prev = money
	}

	return money
}
