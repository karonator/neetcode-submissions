func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0

	var sb strings.Builder
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			sb.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			sb.WriteByte(word2[j])
			j++
		}
	}
	return sb.String()
}
