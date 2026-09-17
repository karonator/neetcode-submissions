func tokenize(path string) []string {
	tokens := make([]string, 0)
	start := 0
	end := start

	for end < len(path) {
		if path[end] == '/' {
			if end - start > 1 {
				tokens = append(tokens, path[start + 1: end])
			}
			start = end 
		}
		end++
	}
	if end - start > 1 {
		tokens = append(tokens, path[start + 1: end])
	}
	return tokens
}

func simplifyPath(path string) string {
	tokens := tokenize(path)
	stack := make([]string, 0)

	for _, token := range tokens {
		if token == "." {
			continue
		} else if token == ".." { 
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
		} else {
			stack = append(stack, token)
		}
	}

	var sb strings.Builder
	if len(stack) == 0 {
		sb.WriteString("/")
	}
	for _, chunk := range stack {
		sb.WriteString("/")
		sb.WriteString(chunk)
	}

	return sb.String()
}
