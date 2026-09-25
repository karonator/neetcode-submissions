func minDistance(word1 string, word2 string) int {
    data := make([][]int, len(word1) + 1)
	for i := range data {
		data[i] = make([]int, len(word2) + 1)
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if i == 0 {
				data[i][j] = j
			} else if j == 0 {
				data[i][j] = i
			}
		}
	}
	for i := 1; i < len(data); i++ {
		for j := 1; j < len(data[0]); j++ {
			a := word1[i - 1]
			b := word2[j - 1]

			if a == b {
				data[i][j] = data[i - 1][j - 1]
			} else {
				// replace
				replace := data[i - 1][j - 1] + 1
				// insert
				insert := data[i][j - 1] + 1
				// delete
				del := data[i - 1][j] + 1
				data[i][j] = min(replace, insert, del)
			}
		}
	}
	return data[len(data) - 1][len(data[0]) - 1]
}
