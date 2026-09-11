func calPoints(operations []string) int {
	stack := []int{}
	ans := 0

	for _, op := range operations {
		if op == "+" {
			stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
			ans += stack[len(stack) - 1]
		} else if op == "D" {
			stack = append(stack, stack[len(stack) - 1] * 2)
			ans += stack[len(stack) - 1]
		} else if op == "C" {
			ans -= stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		} else {
			x, _ := strconv.Atoi(op)
			stack = append(stack, x)
			ans += stack[len(stack) - 1]
		}
	}

	return ans
}
