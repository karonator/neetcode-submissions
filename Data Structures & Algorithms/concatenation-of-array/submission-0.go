func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums) * 2)
	for i := range(nums) {
		ans[i] = nums[i]
		ans[i + len(nums)] = nums[i]
	}
	return ans
}
