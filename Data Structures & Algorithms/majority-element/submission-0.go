func majorityElement(nums []int) int {
    candidate := math.MaxInt32
	votes := 0

	for _, n := range(nums) {
		if votes == 0 {
			candidate = n
			votes++
		} else {
			if candidate == n {
				votes ++
			} else {
				votes --
				if votes == 0 {
					candidate = n
				}
			}
		}
	}
	return candidate
}
