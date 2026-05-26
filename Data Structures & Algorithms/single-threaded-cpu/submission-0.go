type Task struct {
	Enque      int
	Processing int
	Index      int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { 
	if h[i].Processing != h[j].Processing {
		return h[i].Processing < h[j].Processing
	} else {
		return h[i].Index < h[j].Index
	}
}

func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	queue := []Task{}
	for i, task := range(tasks) {
		queue = append(queue, Task {
			Enque:  	task[0],
			Processing: task[1],
			Index: 		i,
		})
	}

	sort.Slice(queue, func(i, j int) bool {
    	return queue[i].Enque < queue[j].Enque
  	})

	ready := &TaskHeap{}
	heap.Init(ready)

	current := 0
	time := 0
	ans := []int{}

	for len(*ready) > 0 || current != len(queue) {
		for current < len(queue) && queue[current].Enque <= time {
			heap.Push(ready, queue[current])
			current++
		}
		if len(*ready) == 0 && current < len(queue) {
			time = queue[current].Enque
			continue
		}

		current_task := heap.Pop(ready).(Task)
		ans = append(ans, current_task.Index)
		time += current_task.Processing
	}

	return ans
}
