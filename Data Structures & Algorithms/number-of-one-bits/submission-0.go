func hammingWeight(n int) int {
	ans := 0
	for i := 0; i <= 31; i++ {
		lastBit := n & 1
		if lastBit == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}
