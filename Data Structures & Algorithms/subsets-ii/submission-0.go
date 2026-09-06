func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

	var backtrack func(int, []int)
	backtrack = func(start int, cur []int) {
		cp := make([]int, len(cur))
		copy(cp, cur)
		ans = append(ans, cp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i - 1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrack(i + 1, cur)
			cur = cur[:len(cur) - 1]
		}
	}

	backtrack(0, []int{})
	return ans
}
