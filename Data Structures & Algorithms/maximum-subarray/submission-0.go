func max (a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
    sums := make([]int, len(nums))
	for i, x := range nums {
		if i == 0 {
			sums[0] = x
		} else {
			if sums[i - 1] > 0 {
				sums[i] = sums[i - 1] + x
			} else {
				sums[i] = x
			}
			maxSum = max(maxSum, sums[i])
		}
	}
	return maxSum
}
