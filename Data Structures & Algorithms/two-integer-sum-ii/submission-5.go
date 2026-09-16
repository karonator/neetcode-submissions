func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for numbers[left] + numbers[right] != target && left != right {
		if numbers[left] + numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{left + 1, right + 1}
}
