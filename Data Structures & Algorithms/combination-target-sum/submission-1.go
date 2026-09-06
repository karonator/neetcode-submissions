func combinationSum(nums []int, target int) [][]int {
    ans := make([][]int, 0)

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

		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i, cur, sum + nums[i])
			cur = cur[:len(cur) - 1]
		}
	}
	backtrack(0, []int{}, 0)
	return ans
}
