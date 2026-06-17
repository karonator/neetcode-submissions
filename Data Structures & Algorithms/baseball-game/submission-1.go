func calPoints(operations []string) int {
	stack := []int{}
	result := 0

	for _, s := range(operations) {
		if s == "+" {
			a := stack[len(stack) - 1]
			b := stack[len(stack) - 2]
			stack = append(stack, a + b)
			result += (a + b)

			continue
		}
		if s == "C" {
			a := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result -= a

			continue
		}
		if s == "D" {
			a := stack[len(stack) - 1]
			stack = append(stack, 2 * a)
			result += 2 * a

			continue
		}
		current, _ := strconv.Atoi(s)
		stack = append(stack, current)
		result += current
	}
	return result
}
