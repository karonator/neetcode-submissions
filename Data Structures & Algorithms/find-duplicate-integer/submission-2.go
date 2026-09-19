func abs (a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findDuplicate(nums []int) int {
    for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			return abs(nums[i])
		} else {
			nums[idx] = -nums[idx]
		}
	}
	return -1
}
