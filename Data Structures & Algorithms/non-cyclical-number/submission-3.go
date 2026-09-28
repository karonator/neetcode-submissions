func isHappy(n int) bool {
	cnt := func(x int) int {
		ans := 0
		for x >= 10 {
			digit := x % 10
			ans += digit * digit
			x -= digit
			x /= 10
		}
		ans += x * x
		return ans
	}

	steps := make(map[int]struct{})

	for n != 1 {
		tmp := cnt(n)
		if _, found := steps[tmp]; found {
			return false
		}
		steps[tmp] = struct{}{}
		n = tmp
	}

	return true
}
