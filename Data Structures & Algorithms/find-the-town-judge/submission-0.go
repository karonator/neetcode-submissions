func findJudge(n int, trust [][]int) int {
	social := make(map[int]int)

	for i := range(trust) {
		social[trust[i][0]] -= 1
		social[trust[i][1]] += 1
	}

	for k, v := range(social) {
		if v == len(social) - 1 {
			return k
		}
	}

	return -1
}
