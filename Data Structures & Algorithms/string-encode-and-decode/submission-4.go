type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var header strings.Builder
	var body strings.Builder

	for i, str := range strs {
		header.WriteString(strconv.Itoa(len(str)))
		if i != len(strs) - 1 {
			header.WriteString(",")
		}
		body.WriteString(str)
	}
	return header.String() + "#" + body.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	header := ""
	body := ""
	for i, c := range encoded {
		if c == '#' {
			header = encoded[:i]
			body = encoded[i+1:]
			break
		}
	}
	headerChunks := strings.Split(header, ",")
	start := 0
	ans := make([]string, len(headerChunks))
	for i := range headerChunks {
		chunkLen, _ := strconv.Atoi(headerChunks[i])
		ans[i] = body[start: start + chunkLen]
		start = start + chunkLen
	}
	return ans
}
