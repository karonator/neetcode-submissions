func longestConsecutive(nums []int) int {
	exists := make(map[int]struct{})
	for _, num := range nums {
		exists[num] = struct{}{}
	}

	starts := make([]int, 0)
	for num := range exists {
		if _, found := exists[num - 1]; !found {
			starts = append(starts, num)
		}
	}

	maxLen := 0
	for _, start := range starts {
		curLen := 1
		cur := start
		for {
			if _, found := exists[cur + 1]; !found {
				break
			}
			cur++
			curLen++
		}
		maxLen = max(maxLen, curLen)
	}

	return maxLen
}
