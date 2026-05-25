func abs_dist(point []int) float64 {
	return math.Sqrt(float64(point[0] * point[0] + point[1] * point[1]))
}

type Elem struct {
	data []int
	dist float64
}

type MyHeap []Elem

func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(Elem))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	hp := &MyHeap{}
	heap.Init(hp)

	for _, point := range(points) {
		dist := abs_dist(point)
		newone := Elem {
			data: point,
			dist: dist,
		}
		if len(*hp) < k {
			heap.Push(hp, newone)
		} else {
			if (*hp)[0].dist > dist {
				heap.Pop(hp)
				heap.Push(hp, newone)
			}
		}
	}

	result := [][]int{}
	for len(*hp) > 0 {
		result = append(result, heap.Pop(hp).(Elem).data)
	}
	return result
}
