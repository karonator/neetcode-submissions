/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, maxDepth int) int {
	if root == nil {
		return 0
	}

	return max(max(core(root.Left, maxDepth + 1), core(root.Right, maxDepth + 1)), maxDepth + 1) 
}

func maxDepth(root *TreeNode) int {
    return core(root, 0)
}
