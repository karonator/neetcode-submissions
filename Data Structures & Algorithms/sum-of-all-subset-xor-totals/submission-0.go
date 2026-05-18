func counter(nums []int, index int, ans int) int {
	if len(nums) == index {
		return ans
	}

	return counter(nums, index + 1, ans) + counter(nums, index + 1, ans ^ nums[index])
}

func subsetXORSum(nums []int) int {
	return counter(nums, 0, 0)
}
