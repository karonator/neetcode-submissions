func findJudge(n int, trust [][]int) int {
	in := make(map[int]int)
	out := make(map[int]int)
	for i := range trust {
		out[trust[i][0]]++
		in[trust[i][1]]++
	}

	for i := range in {
		if in[i] == n - 1 && out[i] == 0 {
			return i
		}
	}

	return -1
}
