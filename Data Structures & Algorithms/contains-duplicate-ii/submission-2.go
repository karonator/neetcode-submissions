func containsNearbyDuplicate(nums []int, k int) bool {
	left := 0
	window := make(map[int]struct{})

	for right := range(nums) {
		if right - left > k {
			delete(window, nums[left])
			left++
		}
		if _, found := window[nums[right]]; found == true {
			return true
		}
		window[nums[right]] = struct{}{}
	}
	fmt.Println(window)
	return false
}
