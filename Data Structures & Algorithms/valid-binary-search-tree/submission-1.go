/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func validate(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	left := validate(root.Left, min, root.Val)
	right := validate(root.Right, root.Val, max)

	return left && right
}

func isValidBST(root *TreeNode) bool {
    return validate(root, -1001, 1001)
}
