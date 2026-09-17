/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var checkTreesSame func(*TreeNode, *TreeNode) bool
	checkTreesSame = func(root1 *TreeNode, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		return root1.Val == root2.Val && checkTreesSame(root1.Left, root2.Left) && checkTreesSame(root1.Right, root2.Right)
	}

	var findAndCheck func(*TreeNode) bool
	findAndCheck = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		return node.Val == subRoot.Val && checkTreesSame(node, subRoot) || findAndCheck(node.Left) || findAndCheck(node.Right)
	}

	return findAndCheck(root)
}
