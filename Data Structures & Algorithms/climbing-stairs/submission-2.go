func climbStairs(n int) int {
	stairs := make([]int, n + 1)
	stairs[0] = 1
	stairs[1] = 1

	for i := 2; i <= n; i++ {
		stairs[i] = stairs[i - 1] + stairs[i - 2]
	}
	return stairs[n]
}
