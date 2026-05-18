func combinationSum(nums []int, target int) [][]int {
	answer := make([][]int, 0)
	
	current := []int{}

	var core func(index int, sum int)
	core = func(index int, sum int) {
		if sum == target {
			cp := make([]int, len(current))
			copy(cp, current)
			answer = append(answer, cp)
			return
		}

		if sum > target || index == len(nums) {
			return
		}

		current = append(current, nums[index])
		core(index, sum + nums[index])
		current = current[0: len(current) - 1]
		core(index + 1, sum)
	}
	core(0, 0)
	return answer
}
