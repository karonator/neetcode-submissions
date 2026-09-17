/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */



func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node)

	dummy := &Node{}
	tail := dummy

	runner := head
	for runner != nil {
		newNode := &Node{
			Val: runner.Val,
		}
		tail.Next = newNode
		tail = tail.Next
		oldToNew[runner] = newNode
		runner = runner.Next
	}

	runner = head
	for runner != nil {
		oldToNew[runner].Random = oldToNew[runner.Random]
		runner = runner.Next
	}

	return dummy.Next

}
