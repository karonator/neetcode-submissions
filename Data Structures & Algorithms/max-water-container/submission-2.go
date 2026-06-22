func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(heights []int) int {
	ans := 0

	i := 0
	j := len(heights) - 1

	for i < j {
		water := min(heights[i], heights[j]) * (j - i)
		if water > ans {
			ans = water
		}

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}
