type Task struct {
	Freq int
	Name byte
}

type IntHeap []Task

func (h IntHeap) Len() int { 
  return len(h) 
}
func (h IntHeap) Less(i, j int) bool { 
  return h[i].Freq > h[j].Freq
}
func (h IntHeap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i] 
}

func (h *IntHeap) Push(x interface{}) {
   *h = append(*h, x.(Task))
}

func (h *IntHeap) Pop() interface{} {
   old := *h
   n := len(old)
   x := old[n-1]
   *h = old[0 : n-1]
   return x
}

func leastInterval(tasks []byte, n int) int {
	freqs := make(map[byte]int)
	for _, task := range tasks {
		freqs[task]++
	}
	tasksHeap := IntHeap{}
	for key, freq := range freqs {
		heap.Push(&tasksHeap, Task{
			Freq: freq,
			Name: key,
		})
	}
	cooldown := make(map[int]Task)

	step := 0
	for tasksHeap.Len() > 0 || len(cooldown) > 0 {
		if task, found := cooldown[step]; found {
			heap.Push(&tasksHeap, task)
			delete(cooldown, step)
		}
		if tasksHeap.Len() > 0 {
			task := heap.Pop(&tasksHeap).(Task)
			if task.Freq > 1 {
				cooldown[step + n + 1] = Task{
					Freq: task.Freq - 1,
					Name: task.Name,
				}
			}
		}
		step++
	}

	return step
}
