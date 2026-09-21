func countComponents(n int, edges [][]int) int {
    nodes := make(map[int][]int)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}

	visited := make(map[int]struct{})
	ans := 0

	for i := range n {
		if _, found := visited[i]; found {
			continue
		}
		ans++
		queue := []int{i}
		for len(queue) > 0 {
			node := queue[len(queue) - 1]
			queue = queue[:len(queue) - 1]
			if _, found := visited[node]; found {
				continue
			}
			visited[node] = struct{}{}
			queue = append(queue, nodes[node]...)
		}
	}

	return ans
}
