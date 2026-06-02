func evalRPN(tokens []string) int {
	mem := []int{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			a := mem[len(mem) - 1]
			mem = mem[:len(mem) - 1]

			b := mem[len(mem) - 1]
			mem = mem[:len(mem) - 1]

			switch tokens[i] {
				case "+":
					mem = append(mem, a + b)
				case "-":
					mem = append(mem, b - a)
				case "*":
					mem = append(mem, a * b)
				case "/":
					mem = append(mem, b / a)
			}
			fmt.Println(mem)

		} else {
			val, _ := strconv.Atoi(tokens[i])
			mem = append(mem, val)
			fmt.Println(mem)
		}
	}

	return mem[len(mem) - 1]
}
