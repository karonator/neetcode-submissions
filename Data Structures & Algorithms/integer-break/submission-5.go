func integerBreak(n int) int {
	// max 3, then max 2
	ans := 1

	if n == 2 { return 1 }
	if n == 3 { return 2 }

	if n % 3 == 0 {
		ans = int(math.Pow(3, float64(n / 3)))
	}
	if n % 3 == 1 {
		ans = int(math.Pow(3, float64((n - 4) / 3))) * 4
	}
	if n % 3 == 2 {
		ans = int(math.Pow(3, float64((n - 2) / 3))) * 2
	}
	return ans
}
