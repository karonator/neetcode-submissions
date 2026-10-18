func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 1
   	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
