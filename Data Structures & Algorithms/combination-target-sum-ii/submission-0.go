func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var backtrack func(int, []int, int)
	backtrack = func(start int, cur []int, sum int) {
		if sum == target {
			cp := make([]int, len(cur))
			copy(cp, cur)
			ans = append(ans, cp)
		}
		if sum >= target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i - 1] {
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(i + 1, cur, sum + candidates[i])
			cur = cur[:len(cur) - 1]
		}
	}
	backtrack(0, []int{}, 0)
	return ans
}
