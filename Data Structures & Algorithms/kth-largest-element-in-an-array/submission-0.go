type MyHeap []int

func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	heapy := &MyHeap{}
	heap.Init(heapy)

	for _, num := range(nums) {
		if len(*heapy) < k {
			heap.Push(heapy, num)
		} else {
			if (*heapy)[0] < num {
				heap.Pop(heapy)
				heap.Push(heapy, num)
			}
		}
	}

	return heap.Pop(heapy).(int)
}
