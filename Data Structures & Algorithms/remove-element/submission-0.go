func removeElement(nums []int, val int) int {
    fast := 0
	unique := 0
	for slow := 0; slow < len(nums); slow ++ {
		if nums[slow] == val {
			if fast < slow {
				fast = slow
			}
			for fast < len(nums) {
				if nums[fast] != val {
					nums[slow], nums[fast] = nums[fast], nums[slow]
					unique++
					break
				}
				fast ++
			}
		} else {
			unique ++
		}
	}
	return unique
}
