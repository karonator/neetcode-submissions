func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	for i := range nums {
		prevPrefix := 1
		if i > 0 {
			prevPrefix = prefix[i - 1]		
		}
		prefix[i] = prevPrefix * nums[i]
	}
	
	for i := len(nums) - 1; i >= 0; i-- {
		prevSuffix := 1
		if i < len(nums) - 1 {
			prevSuffix = suffix[i + 1]		
		}
		suffix[i] = prevSuffix * nums[i]
	}

	for i := range nums {
		pref := 1
		if i > 0 {
			pref = prefix[i - 1]
		}
		suff := 1
		if i < len(nums) - 1 {
			suff = suffix[i + 1]
		}
		result[i] = pref * suff
	}
	return result
}