/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func clone(node *Node, cache map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }
    current, found := cache[node]
    if found {
        return current
    }
    current = &Node {
        Val: node.Val,
    }
    cache[node] = current

    for i := range(node.Neighbors) {
        child := clone(node.Neighbors[i], cache)
        current.Neighbors = append(current.Neighbors, child)
    }
    return current
}

func cloneGraph(node *Node) *Node {
    nodes_mapper := make(map[*Node]*Node)
    return clone(node, nodes_mapper)
}
