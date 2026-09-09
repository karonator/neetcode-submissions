func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		if i == 0 {
			left[i] = 0
			right[len(height) - 1 - i] = 0
		} else {
			left[i] = max(left[i - 1], height[i - 1])
			right[len(height) - 1 - i] = max(right[len(height) - i], height[len(height) - i])
		}
	}

	ans := 0
	for i := range height {
		capacity := min(right[i], left[i])
		water := max(0, capacity - height[i])
		ans += water
	}
	return ans
}
