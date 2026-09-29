type Pair struct {
	Num int
	Data string
}

func replicate(s string, n int) string {
	ans := ""
	for range n {
		ans += s
	}
	return ans
}

func decodeString(s string) string {
	stack := []Pair{}

	num := 0
	cur := ""

	for i := range s {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = 10 * num + int(c - '0')
		} else if c == '[' {
			stack = append(stack, Pair{
				Num: num,
				Data: cur,
			})
			num = 0
			cur = ""
		} else if c == ']' {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			cur = top.Data + replicate(cur, top.Num)  
		} else {
			cur += string(c)
		}
	}

	return cur
}
