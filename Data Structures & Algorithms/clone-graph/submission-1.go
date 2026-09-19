/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	registry := make(map[*Node]*Node)

	var clone func(*Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		newNode := &Node{
			Val: node.Val,
		}
		registry[node] = newNode
	
		newNeigbours := []*Node{}
		for i := range node.Neighbors {
			if neiNode, found := registry[node.Neighbors[i]]; !found {
				newNeigbours = append(newNeigbours, clone(node.Neighbors[i]))
			} else {
				newNeigbours = append(newNeigbours, neiNode)
			}
		}
		newNode.Neighbors = newNeigbours
		return newNode
	}
	return clone(node)
}
