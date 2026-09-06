func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	
	var backtrack func(int, []int)
	backtrack = func(i int, cur []int) {
		if i == len(nums) {
			cp := make([]int, len(cur))
			copy(cp, cur)
			result = append(result, cp)
			return
		}

		cur = append(cur, nums[i])
		backtrack(i + 1, cur)

		cur = cur[:len(cur) - 1]
		backtrack(i + 1, cur)
	}
	backtrack(0, []int{})

	return result
}
