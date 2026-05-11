func minSubArrayLen(target int, nums []int) int {
	left := 0
	current_sum := 0
	answer := len(nums) + 1
	
	for right := 0; right < len(nums); right++ {
		current_sum += nums[right]

		for current_sum >= target {
			if right - left + 1 < answer {
				answer = right - left + 1
			}

			current_sum -= nums[left]
			left ++
		}
	}

	if answer == len(nums) + 1 {
		return 0
	}

	return answer
}
