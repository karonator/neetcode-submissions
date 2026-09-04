type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
    Data *MinIntHeap
	K 	 int
}


func Constructor(k int, nums []int) KthLargest {
	ans := KthLargest{
		Data: &MinIntHeap{},
		K: k,
	}
	heap.Init(ans.Data)
	for _, num := range nums {
		ans.Add(num)
	}
	return ans
}


func (this *KthLargest) Add(val int) int {
	if len(*this.Data) < this.K {
		heap.Push(this.Data, val)
		return (*this.Data)[0]
	} else {
		minElem := (*this.Data)[0]
		if val > minElem {
			heap.Pop(this.Data)
			heap.Push(this.Data, val)
		}
	}
	return (*this.Data)[0]
}
