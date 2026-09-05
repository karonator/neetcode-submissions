func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, ast := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, ast)
		} else {
			last := stack[len(stack) - 1]
			dead := false
			for last > 0 && ast < 0 && len(stack) > 0 {
				if abs(last) > abs(ast) {
					dead = true
					break
				} else if abs(last) == abs(ast) {
					stack = stack[:len(stack) - 1]
					dead = true
					break
				} else {
					stack = stack[:len(stack) - 1]
					if len(stack) > 0 {
						last = stack[len(stack) - 1]
					}
				}
			}
			if !dead {
				stack = append(stack, ast)
			}
		}
	}
	return stack
}
