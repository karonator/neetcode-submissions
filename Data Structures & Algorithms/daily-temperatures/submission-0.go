func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 {
			last := stack[len(stack) - 1]
			lastVal := temperatures[last]
			if lastVal >= temp {
				break
			}
			ans[last] = i - last
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return ans
}
