type node struct {
	coord [2]int
	dist int
}

type minHeap []*node

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(val interface{}) {
	*h = append(*h, val.(*node))
}

func (h *minHeap) Pop() interface{} {
	heapDerefrenced := *h

	size := len(heapDerefrenced)
	val := heapDerefrenced[size-1]
	*h = heapDerefrenced[:size-1]

	return val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumEffortPath(heights [][]int) int {
	visited := make(map[[2]int]struct{})
	dists := make(map[[2]int]int)
	dists[[2]int{0, 0}] = 0

	queue := minHeap{
		&node{
			coord: [2]int{0, 0},
			dist: 0,
		},
	}

	offsets := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for queue.Len() > 0 {
		nodeData := heap.Pop(&queue).(*node)

		if _, found := visited[nodeData.coord]; found {
			continue
		}
		visited[nodeData.coord] = struct{}{}

		for _, offset :=  range offsets {
			x := nodeData.coord[0] + offset[0]
			y := nodeData.coord[1] + offset[1]
			key := [2]int{x, y}

			if x >= 0 && y >= 0 && x < len(heights) && y < len(heights[0]) {
				newDist := max(nodeData.dist, abs(heights[nodeData.coord[0]][nodeData.coord[1]] - heights[x][y]))
				if dist, found := dists[key]; found {
					if newDist < dist {
						dists[key] = newDist
						heap.Push(&queue, &node{
							coord: [2]int{x, y},
							dist: newDist,
						})
					}
				} else {
					dists[key] = newDist
					heap.Push(&queue, &node{
						coord: key,
						dist: newDist,
					})
				}
			}
		}
	}
	ans, _ := dists[[2]int{len(heights) - 1, len(heights[0]) - 1}]
	return ans
}
