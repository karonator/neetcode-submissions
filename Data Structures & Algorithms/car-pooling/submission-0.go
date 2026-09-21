import (
	"slices"
	"cmp"
)

type CarChunkData struct {
	people int
	to     int
}

type MyHeap []*CarChunkData

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i].to < h[j].to
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(val interface{}) {
	*h = append(*h, val.(*CarChunkData))
}

func (h *MyHeap) Pop() interface{} {
	heapDerefrenced := *h

	size := len(heapDerefrenced)
	val := heapDerefrenced[size-1]
	*h = heapDerefrenced[:size-1]

	return val
}

func carPooling(trips [][]int, capacity int) bool {
	slices.SortFunc(trips, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	free := capacity
	outQueue := &MyHeap{}
	for _, trip := range trips {
		cur := trip[1]
		for outQueue.Len() > 0 && (*outQueue)[0].to <= cur {
			chunk := heap.Pop(outQueue).(*CarChunkData)
			free += chunk.people
		}
		if free < trip[0] {
			return false
		} else {
			free -= trip[0]
			heap.Push(outQueue, &CarChunkData{
				people: trip[0],
				to: trip[2],
			})
		}
	}
	return true
}
