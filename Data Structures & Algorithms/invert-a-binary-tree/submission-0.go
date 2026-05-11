/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(node *TreeNode) {
	if node == nil {
		return
	}

	tempNode := node.Right
	node.Right = node.Left
	node.Left = tempNode

	core(node.Left)
	core(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
    core(root)
	return root
}
