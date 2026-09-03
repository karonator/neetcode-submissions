/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var insert func(*TreeNode) *TreeNode
	insert = func(node *TreeNode) *TreeNode {
		if node == nil {
			return &TreeNode {
				Val: val,
			}
		}
		if val > node.Val {
			node.Right = insert(node.Right)
		} else {
			node.Left = insert(node.Left)
		}
		return node
	}
	return insert(root)
}
