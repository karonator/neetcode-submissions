func removeDuplicates(nums []int) int {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[slow] == nums[fast] {
			continue
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
