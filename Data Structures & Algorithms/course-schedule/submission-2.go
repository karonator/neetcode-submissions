func canFinish(numCourses int, prerequisites [][]int) bool {
    nodes := make(map[int]int)
	children := make(map[int][]int)

	if len(prerequisites) == 0 {
		return true
	}

	for _, v := range(prerequisites) {
		_, start_found := nodes[v[0]]
		if !start_found {
			nodes[v[0]] = 0
		}
		_, end_found := nodes[v[1]]
		if !end_found {
			nodes[v[1]] = 1
		} else {
			nodes[v[1]]++
		}
		children[v[0]] = append(children[v[0]], v[1])
	}

	queue := []int{}

	deleted := 0
	for node, val := range(nodes) {
		if val == 0 {
			queue = append(queue, node)
			delete(nodes, node)
			deleted++
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		for _, child := range(children[node]) {
			nodes[child]--
			if nodes[child] == 0 {
				queue = append(queue, child)
				delete(nodes, child)
				deleted++
			}
		}
		queue = queue[1:]
	}
	
	return deleted == numCourses || len(queue) == 0 && len(nodes) == 0
}
