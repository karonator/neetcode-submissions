/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    mapper := make(map[*Node]*Node)

	runner := head
	var ans_head *Node
	var ans_tail *Node

	for runner != nil {
		node := &Node {
			Val: runner.Val,
		}
		mapper[runner] = node

		if ans_head == nil {
			ans_head = node
			ans_tail = node
		} else {
			ans_tail.Next = node
			ans_tail = ans_tail.Next
		}
		runner = runner.Next
	}

	runner = head
	for runner != nil {
		mapper[runner].Random = mapper[runner.Random]
		runner = runner.Next
	}

	return ans_head
}
