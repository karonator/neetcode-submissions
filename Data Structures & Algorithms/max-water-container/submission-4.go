func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	ans := 0

	for right > left {
		ans = max(ans, (right - left) * min(heights[right], heights[left]))
		if heights[right] < heights[left] {
			right--
		} else {
			left++
		}
	}
	return ans
}
