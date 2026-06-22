func removeDuplicates(nums []int) int {
	last := -999
	k := 0
	
	for i := range(nums) {
		if nums[i] != last {
			last = nums[i]
			nums[k] = nums[i]
			k ++
		}
	}

	return k
}
