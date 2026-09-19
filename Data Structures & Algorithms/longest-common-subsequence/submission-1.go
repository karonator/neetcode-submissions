func longestCommonSubsequence(text1 string, text2 string) int {
	data := make([][]int, len(text1) + 1)
	for i := range data {
		data[i] = make([]int, len(text2) + 1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i - 1] == text2[j - 1] {
				data[i][j] = data[i - 1][j - 1] + 1
			} else {
				data[i][j] = max(data[i - 1][j], data[i][j - 1])
			}
		}
	}
	return data[len(data) - 1][len(data[0]) - 1]
}
