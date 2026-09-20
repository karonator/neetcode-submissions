func plusOne(digits []int) []int {
    transfer := false
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			transfer = true
		} else {
			transfer = false
			break
		}
	}
	if transfer {
		return append([]int{1}, digits...)
	}
	return digits
}
