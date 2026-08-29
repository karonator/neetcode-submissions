func containsNearbyDuplicate(nums []int, k int) bool {
	left := 0
	window := make(map[int]int)
	window[nums[0]]++

	for right := 1; right < len(nums); right ++ {
		window[nums[right]] ++

		if right - left > k {
			window[nums[left]] --
			left ++
		}

		if window[nums[right]] > 1 {
			return true
		}
	}
	return false
}
