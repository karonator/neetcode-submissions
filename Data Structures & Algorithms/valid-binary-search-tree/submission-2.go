/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return core(root.Left, min, root.Val) && core(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
    return core(root, math.MinInt, math.MaxInt)
}
