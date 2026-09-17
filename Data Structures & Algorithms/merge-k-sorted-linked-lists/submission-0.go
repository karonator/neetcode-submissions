/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := MinHeap{}

	for _, listHead := range lists {
		if listHead != nil {
			minHeap = append(minHeap, listHead)
		}
	}

	heap.Init(&minHeap)

	dummyHead := &ListNode{}
	tail := dummyHead

	for minHeap.Len() != 0 {
		node := heap.Pop(&minHeap).(*ListNode)

		next := node.Next
		tail.Next = node
		tail = tail.Next

		if next != nil {
			heap.Push(&minHeap, next)
		}
	}

	return dummyHead.Next
}
