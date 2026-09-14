func topKFrequent(nums []int, k int) []int {
	intToFreq := make(map[int]int)
	for i := range nums {
		intToFreq[nums[i]]++
	}

	freqs := make(map[int][]int)
	for i, freq := range intToFreq {
		if _, found := freqs[freq]; !found {
			freqs[freq] = make([]int, 0)
		}
		freqs[freq] = append(freqs[freq], i)
	}
	fmt.Println(freqs)
	ans := make([]int, 0)
	count := 0
	for i := len(nums); i >= 0; i-- {
		if intsByFreq, found := freqs[i]; found {
			for _, j := range intsByFreq {
				ans = append(ans, j)
				count++
				if count == k {
					return ans
				}
			}
		}
	}
	return ans
}
