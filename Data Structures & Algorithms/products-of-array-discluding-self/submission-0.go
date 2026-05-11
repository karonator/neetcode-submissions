func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		tmp := 1
		for j := range nums {
			if i != j {
				tmp *= nums[j]
			}
		}
		result[i] = tmp
	}

	return result
}
