func collide(a int, b int) bool {
	return a > 0 && b < 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func asteroidCollision(asteroids []int) []int {
	result := []int{}

	for _, asteroid := range asteroids {
		alive := true

		for alive && len(result) > 0 && collide(result[len(result)-1], asteroid) {
			last := result[len(result)-1]

			if abs(last) < abs(asteroid) {
				// верхушка стека погибает
				result = result[:len(result)-1]
			} else if abs(last) == abs(asteroid) {
				// погибают оба
				result = result[:len(result)-1]
				alive = false
			} else {
				// текущий астероид погибает
				alive = false
			}
		}

		if alive {
			result = append(result, asteroid)
		}
	}

	return result
}