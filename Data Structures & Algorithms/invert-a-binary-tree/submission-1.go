/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	core(root.Left)
	core(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	core(root)
	return root
}
