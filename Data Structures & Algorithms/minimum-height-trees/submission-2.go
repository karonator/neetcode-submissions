func findMinHeightTrees(n int, edges [][]int) []int {
	connections := make(map[int]map[int]struct{})

	for i := range n {
		connections[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		start := edge[0]
		end := edge[1]

		connections[start][end] = struct{}{}
		connections[end][start] = struct{}{}
	}

	queue := []int{}
	for nodeID := range connections {
		if len(connections[nodeID]) <= 1 {
			queue = append(queue, nodeID)
		}
	}

	for len(connections) > 2 {
		newQueue := []int{}
		for i := range queue {
			for k := range connections[queue[i]] {
				delete(connections[k], queue[i])
				if len(connections[k]) == 1 {
					newQueue = append(newQueue, k)
				}
			}
		}
		for i := range queue {
			delete(connections, queue[i])
		}
		queue = newQueue
	}
	return queue
}
