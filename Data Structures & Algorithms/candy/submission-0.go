func candy(ratings []int) int {
	left := make([]int, len(ratings))
	for i := range left {
		if i == 0 {
			left[0] = 1
			continue
		}
		if ratings[i] > ratings[i - 1] {
			left[i] = left[i - 1] + 1
		} else {
			left[i] = 1
		}
	}

	right := make([]int, len(ratings))
	for i := len(ratings) -1; i >= 0; i-- {
		if i == len(ratings) -1 {
			right[len(ratings) -1] = 1
			continue
		}
		if ratings[i] > ratings[i + 1] {
			right[i] = right[i + 1] + 1
		} else {
			right[i] = 1
		}
	}

	ans := 0
	for i := range len(ratings) {
		ans += max(left[i], right[i])
	}

	return ans
}
