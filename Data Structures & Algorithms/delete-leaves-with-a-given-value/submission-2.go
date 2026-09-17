/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteCascaded(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = deleteCascaded(root.Left, target)
	root.Right = deleteCascaded(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    return deleteCascaded(root, target)
}
