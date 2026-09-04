type Point struct {
	Coords 	[]int
	Dist 	int
}

type MaxHeap []Point

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Dist > h[j].Dist }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dist(point []int) int {
	return point[0] * point[0] + point[1] * point[1]
}

func kClosest(points [][]int, k int) [][]int {
	data := &MaxHeap{}
	heap.Init(data)

	for _, point := range points {
		newPoint := Point{
			Coords: point,
			Dist: dist(point),
		}
		if len(*data) < k {
			heap.Push(data, newPoint)
		} else {
			maxDistPoint := (*data)[0]
			if maxDistPoint.Dist > newPoint.Dist {
				heap.Pop(data)
				heap.Push(data, newPoint)
			}
		}
	}

	result := make([][]int, k)
	i := 0
	for len(*data) > 0 {
		result[i] = (heap.Pop(data).(Point)).Coords
		i++
	}
	
	return result
}
