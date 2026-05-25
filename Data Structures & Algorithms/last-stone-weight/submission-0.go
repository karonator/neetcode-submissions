type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func lastStoneWeight(stones []int) int {
	h := (*IntHeap)(&stones)
	heap.Init(h)

	for len(*h) > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		if a == b {
			continue
		} else {
			heap.Push(h, abs(a - b))
		}
	}
	if len(*h) == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
