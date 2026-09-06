func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	post := make([]int, len(nums))
	ans := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			pre[i] = nums[i]
			post[len(nums) - 1] = nums[len(nums) - 1]
		} else {
			pre[i] = pre[i - 1] * nums[i]
			post[len(nums) - 1 - i] = post[len(nums) - i] * nums[len(nums) - 1 - i]
		}
	}

	for i := range nums {
		ans[i] = 1
		if i > 0 {
			ans[i] *= pre[i - 1]
		}
		if i < len(nums) - 1 {
			ans[i] *= post[i + 1]
		}
	}

	return ans
}
