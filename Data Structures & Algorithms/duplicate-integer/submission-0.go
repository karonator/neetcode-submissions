func hasDuplicate(nums []int) bool {
	cache := make(map[int]struct{})
	for i := range(nums) {
		if _, found := cache[nums[i]]; found {
			return true
		}
		cache[nums[i]] = struct{}{}
	}
	return false
}
