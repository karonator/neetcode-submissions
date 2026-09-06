func combine(n int, k int) [][]int {
	ans := make([][]int, 0)

	var backtrack func(int, []int)
	backtrack = func(start int, cur []int) {
		if len(cur) == k {
			cp := make([]int, len(cur))
			copy(cp, cur)
			ans = append(ans, cp)
			return
		}
		for i := start; i <= n; i++ {
			cur = append(cur, i)
			backtrack(i + 1, cur)
			cur = cur[:len(cur) - 1]
		}
	}
	backtrack(1, []int{})
	return ans
}
