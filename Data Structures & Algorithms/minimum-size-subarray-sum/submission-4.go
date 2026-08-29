
func minSubArrayLen(target int, nums []int) int {
	ans := 0
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if sum >= target {
			for sum - nums[left] >= target {
				sum -= nums[left]
				left++
			}
			if ans > 0 {
				ans = min(ans, right - left + 1)
			} else {
				ans = right - left + 1
			}
		}
	}
	return ans
}
