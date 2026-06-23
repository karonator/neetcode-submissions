func reverse(nums []int, start int, end int) {
	i := start
	j := end

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}


func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, len(nums) - 1)
}
