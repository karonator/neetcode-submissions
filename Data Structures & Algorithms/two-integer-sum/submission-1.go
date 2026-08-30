func twoSum(nums []int, target int) []int {
	elems := make(map[int]int)
	for i, e := range(nums) {
		elems[e] = i
	}

	for i, e := range(nums) {
		diff := target - e
		if j, found := elems[diff]; found && i != j {
			if i < j {
				return []int{i, j}
			} else {
				return []int{j, i}
			}
		}
	}

	return nil
}
