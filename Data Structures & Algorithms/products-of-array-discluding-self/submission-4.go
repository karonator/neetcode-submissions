func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	// prefix
	for i := range nums {
		if i == 0 {
			ans[i] = nums[i]
		} else {
			ans[i] = ans[i - 1] * nums[i]
		}
	}

	// suffix
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i > 0 {
			ans[i] = suffix * ans[i - 1]
		} else {
			ans[i] = suffix
		}
		suffix *= nums[i]
	}

	return ans
}
