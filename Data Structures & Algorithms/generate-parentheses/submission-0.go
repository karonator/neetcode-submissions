func generateParenthesis(n int) []string {
	ans := []string{}
	current := ""

	open := 0
	closed := 0

	var backtrack func()
	backtrack = func() {
		if open == n && closed == n {
			ans = append(ans, current)
			return
		}

		if open <= n {
			current += "("
			open++
			backtrack()
			open--
			current = current[:len(current) - 1]
		}
		if closed <= n && closed < open {
			current += ")"
			closed++
			backtrack()
			closed--
			current = current[:len(current) - 1]
		}
	}

	backtrack()
	return ans
}
