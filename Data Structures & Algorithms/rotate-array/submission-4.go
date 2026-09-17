func rotate(nums []int, k int) {
	k = k % len(nums)

	_rotate := func(start int, end int) {
		for i, j := start, end; i < j; i, j = i + 1, j - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	_rotate(0, len(nums) - 1)
	_rotate(0, k - 1)
	_rotate(k, len(nums) - 1)
}
