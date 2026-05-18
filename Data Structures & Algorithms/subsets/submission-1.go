func subsets(nums []int) [][]int {
	var answer [][]int
	current := []int{}

	var core func(index int)
	core = func(index int) {
		if index == len(nums) {
			cp := make([]int, len(current))
			copy(cp, current)
			answer = append(answer, cp)
			return
		}
		current = append(current, nums[index])
		core(index + 1)
		current = current[:len(current)-1]
		core(index + 1)
	}
	core(0)
	return answer
}
