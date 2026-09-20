func canPartition(nums []int) bool {
    sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum % 2 == 1 {
		return false
	}

	available := make(map[int]struct{})
	available[0] = struct{}{}

	for _, num := range nums {
		add := make([]int, 0)
		for key := range available {
			add = append(add, key + num)
		}
		for _, n := range add {
			available[n] = struct{}{}
		}
		if _, found := available[sum/2]; found {
			return true
		}
	}

	if _, found := available[sum/2]; found {
		return true
	}

	return false
}
