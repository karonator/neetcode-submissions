func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if token == "+" {
			a := stack[len(stack) - 1]
			b := stack[len(stack) - 2]
			stack[len(stack) - 2] = a + b
			stack = stack[:len(stack) - 1]
		} else if token == "-" {
			a := stack[len(stack) - 1]
			b := stack[len(stack) - 2]
			stack[len(stack) - 2] = b - a
			stack = stack[:len(stack) - 1]
		} else if token == "*" {
			a := stack[len(stack) - 1]
			b := stack[len(stack) - 2]
			stack[len(stack) - 2] = a * b
			stack = stack[:len(stack) - 1]
		} else if token == "/" {
			a := stack[len(stack) - 1]
			b := stack[len(stack) - 2]
			stack[len(stack) - 2] = b / a
			stack = stack[:len(stack) - 1]
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}