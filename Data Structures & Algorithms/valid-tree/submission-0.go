func validTree(n int, edges [][]int) bool {
	parents := make([]int, n)
	for i := range n {
		parents[i] = i
	}

	find := func(x int) int {
		chain := []int{x}
		for parents[x] != x {
			x = parents[x]
			chain = append(chain, x)
		}
		for _, i := range chain {
			parents[i] = x
		}
		return x
	}

	union := func(a int, b int) bool {
		ar, br := find(a), find(b)
		if ar == br {
			return false
		}
		parents[ar] = br
		return true
	}

	fmt.Println(parents)

	components := n
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
		components--
	}

	if components > 1 {
		return false
	}

	return true
}
