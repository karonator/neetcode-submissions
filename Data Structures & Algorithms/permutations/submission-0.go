func permute(nums []int) [][]int {
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

		for i := range nums {
			if !used[i] {
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
