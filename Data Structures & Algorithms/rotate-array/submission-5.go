func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse := func(start int, end int) {
		for i, j := start, end; i < j; i, j = i + 1, j - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(0, len(nums) - 1)
	reverse(0, k - 1)
	reverse(k, len(nums) - 1)
}
