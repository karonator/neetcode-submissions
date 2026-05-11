type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lengths := make([]string, 0)
	var data strings.Builder
	for _, s := range strs {
		lengths = append(lengths, strconv.Itoa(len(s)))
		data.WriteString(s)
	}

	return strings.Join(lengths, ",") + "|" + data.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	header_delimiter_index := strings.Index(encoded, "|")
	lengths_str := encoded[0: header_delimiter_index]
	lengths := strings.Split(lengths_str, ",")

	start := header_delimiter_index + 1
	answer := make([]string, 0)
	for _, string_len_str := range lengths {
		string_len, _ := strconv.Atoi(string_len_str)
		answer = append(answer, encoded[start: start + string_len])
		start += string_len
	}
	return answer
}
