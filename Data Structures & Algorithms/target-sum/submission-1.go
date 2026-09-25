func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if target > total || target < -total {
		return 0
	}

	prev := make([]int, 2*total+1)
	offset := total
	prev[offset] = 1

	for _, num := range nums {
		cur := make([]int, 2*total+1)

		for sum := -total; sum <= total; sum++ {
			idx := offset + sum

			if sum-num >= -total {
				cur[idx] += prev[offset+sum-num]
			}

			if sum+num <= total {
				cur[idx] += prev[offset+sum+num]
			}
		}

		prev = cur
	}

	return prev[offset+target]
}