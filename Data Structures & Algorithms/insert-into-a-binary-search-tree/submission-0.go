/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		new_node := &TreeNode{
			Val: val,
		}
		return new_node
	}
	if node.Val > val {
		node.Left = insert(node.Left, val)
	} else {
		node.Right = insert(node.Right, val)
	}
	return node
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}
