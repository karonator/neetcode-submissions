func rangeBitwiseAnd(left int, right int) int {
	cnt := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		cnt++
	}
	left = left << cnt
	return left
}
