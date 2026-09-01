func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	// важно, мы делаем len(nums) чтобы обработать случай
	// когда искомый элемент должен будет быть добавлен в самый конец
	for lo < hi {
		mid := lo + (hi - lo) / 2
		if target <= nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
