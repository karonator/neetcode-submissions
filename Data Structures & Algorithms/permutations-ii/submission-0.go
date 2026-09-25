func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, current)
			ans = append(ans, tmp)
			return
		}

		usedOnThisLevel := make(map[int]struct{})
		for i := range nums {
			_, skip := usedOnThisLevel[nums[i]]
			if !used[i] && !skip {
				usedOnThisLevel[nums[i]] = struct{}{}
				used[i] = true
				current = append(current, nums[i])
				backtrack()
				used[i] = false
				current = current[:len(current) - 1]
			}
		}		
	}

	backtrack()
	return ans
}
