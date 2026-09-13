type Node struct {
	Key 	int
	Dist 	int
}

type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i].Dist < h[j].Dist
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	tree := make(map[int]map[int]int)
	for _,time := range times {
		_, found := tree[time[0]]
		if !found {
			tree[time[0]] = make(map[int]int)	
		}
		tree[time[0]][time[1]] = time[2]
	}

	distHeap := MinHeap{}

	processed := make(map[int]struct{})
	dists := make(map[int]int)
	dists[k] = 0
	heap.Push(&distHeap, Node{
		Key: k,
		Dist: 0,
	})

	getSource := func() int {
		for distHeap.Len() > 0 {
			node := heap.Pop(&distHeap).(Node)
			if _, found := processed[node.Key]; found {
				continue
			} else {
				return node.Key
			}
		}
		return -1
	}

	source := getSource()
	for source > 0 {
		neighbours := tree[source]
		for neighbourID, neighbourPrice := range neighbours {
			newDist := dists[source] + neighbourPrice
			if d, ok := dists[neighbourID]; !ok || newDist < d {
				dists[neighbourID] = newDist
				heap.Push(&distHeap, Node{
					Key: neighbourID,
					Dist: newDist,
				})
			}
		}
		processed[source] = struct{}{}
		source = getSource()
	}

	if len(dists) < n {
		return -1
	}

	ans := 0
	for _, dist := range dists {
		ans = max(ans, dist)
	}

	return ans
}
