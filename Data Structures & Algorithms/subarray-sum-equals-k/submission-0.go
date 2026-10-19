func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	sums[0] = 1

	ans := 0
	prefix := 0
	for _, num := range nums {
		prefix += num
		if amount, found := sums[prefix - k]; found {
			ans += amount
		}
		sums[prefix]++
	}

	return ans
}
