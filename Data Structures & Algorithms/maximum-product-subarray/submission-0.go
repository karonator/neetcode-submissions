func maxProduct(nums []int) int {
    minPr := nums[0]
	maxPr := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		tmp1 := minPr * nums[i]
		tmp2 := maxPr * nums[i]
		minPr = min(nums[i], tmp1, tmp2)
		maxPr = max(nums[i], tmp1, tmp2)
		ans = max(ans, maxPr)
	}

	return ans
}
