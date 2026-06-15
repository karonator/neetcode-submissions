func subsets(nums []int) [][]int {
	results := [][]int{}
	path := []int{}

	var core func(pos int)
	core = func (pos int) {
		path = append(path, nums[pos])
		if len(nums) - 1 == pos {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
		} else {
			core(pos + 1)
		}

		path = path[: len(path) - 1]
		if len(nums) - 1 == pos {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
		} else {
			core(pos + 1)
		}
	}

	core(0)
	return results
}
