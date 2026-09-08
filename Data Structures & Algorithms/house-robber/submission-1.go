func rob(nums []int) int {
    data := make([]int, len(nums))
	data[0] = nums[0]
	if len(nums) > 1 {
		data[1] = max(data[0], nums[1])
	}
	for i := 2; i < len(nums); i++ {
		data[i] = max(data[i - 1], data[i - 2] + nums[i])
	}
	return data[len(nums) - 1]
}
