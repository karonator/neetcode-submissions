func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
    parents := make([]int, n)
	
	for i := range parents { parents[i] = i }

	find := func(a int) int {
		path := []int{a}
		for parents[a] != a {
			path = append(path, parents[a])
			a = parents[a]
		}
		return a
	}

	union := func(a int, b int) bool {
		ra, rb := find(a), find(b)
		if ra != rb {
			parents[ra] = rb
			return true
		}
		return false
	}

	for _, edge := range edges {
		if !union(edge[0] - 1, edge[1] - 1) {
			return edge
		}
	}
	return []int{}
}
