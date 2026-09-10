func sortColors(nums []int) {
	nuls := 0
	ones := 0
	twos := len(nums) - 1

	for ones <= twos {
		if nums[ones] == 0 {
			nums[nuls], nums[ones] = nums[ones], nums[nuls]
			ones++
			nuls++
		} else if nums[ones] == 1 {
			ones++
		} else {
			nums[twos], nums[ones] = nums[ones], nums[twos]
			twos--
		}
	}
}
