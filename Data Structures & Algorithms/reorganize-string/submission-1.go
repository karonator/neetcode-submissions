type MyChar struct {
	c 	 rune
	freq int
}

type MyHeap []MyChar

func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(MyChar))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reorganizeString(s string) string {
	freqs := make(map[rune]int)
	for _, c := range(s) {
		freqs[c]++
	}

	freq_heap := &MyHeap{}
	heap.Init(freq_heap)

	for k, v := range(freqs) {
		heap.Push(freq_heap, MyChar{
			c: k,
			freq: v,
		})
	}

	var buffer bytes.Buffer
	for len(*freq_heap) > 1 {
        a := heap.Pop(freq_heap).(MyChar)
		b := heap.Pop(freq_heap).(MyChar)

		buffer.WriteRune(a.c)
		buffer.WriteRune(b.c)

		a.freq--
		b.freq--
		
		if a.freq > 0 {
			heap.Push(freq_heap, a)
		}
		if b.freq > 0 {
			heap.Push(freq_heap, b)
		}
    }

	if len(*freq_heap) > 0 {
		x := heap.Pop(freq_heap).(MyChar)
		if x.freq > 1 {
			return ""
		}
		buffer.WriteRune(x.c)
	}

	return buffer.String()
}
