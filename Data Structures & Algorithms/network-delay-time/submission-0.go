type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	INF := 999999

	tree := make(map[int]map[int]int)
	for _,time := range times {
		_, found := tree[time[0]]
		if !found {
			tree[time[0]] = make(map[int]int)	
		}
		tree[time[0]][time[1]] = time[2]
	}

	processed := make(map[int]struct{})
	dists := make(map[int]int)
	for i := range n {
		if i + 1 == k {
			dists[i + 1] = 0
		} else {
			dists[i + 1] = INF
		}
	}

	getSource := func() int {
		minDist := INF
		ans := -1
		for i, dist := range dists {
			if _, found := processed[i]; found {
				continue
			} 
			if dist < minDist {
				ans = i
				minDist = dist
			}
		}
		return ans
	}

	source := getSource()
	for source > 0 {
		neighbours := tree[source]
		for neighbourID, neighbourPrice := range neighbours {
			newDist := dists[source] + neighbourPrice
			dists[neighbourID] = min(dists[neighbourID], newDist)
		}
		processed[source] = struct{}{}
		source = getSource()
	}

	ans := 0
	for _, dist := range dists {
		if dist == INF {
			return -1
		}
		ans = max(ans, dist)
	}

	return ans
}
