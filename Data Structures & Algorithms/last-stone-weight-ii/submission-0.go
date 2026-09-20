func lastStoneWeightII(stones []int) int {
	variants := make(map[int]struct{})
	variants[0] = struct{}{}

	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	maxVariant := 0
	for i := 0; i < len(stones); i++ {
		add := []int{}
		for key := range variants {
			if key + stones[i] <= sum / 2 {
				add = append(add, key + stones[i])
			}
		}
		for _, n := range add {
			variants[n] = struct{}{}
			maxVariant = max(maxVariant, n)
		}
	}
	return sum - 2 * maxVariant
}
