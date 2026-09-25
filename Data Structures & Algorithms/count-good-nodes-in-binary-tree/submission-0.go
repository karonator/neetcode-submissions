/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var traverse func(*TreeNode, int) int
	traverse = func(node *TreeNode, maximum int) int {
		if node == nil {
			return 0
		}
		ans := 0
		if node.Val >= maximum {
			ans = 1
		}
		ans += traverse(node.Left, max(maximum, node.Val))
		ans += traverse(node.Right, max(maximum, node.Val))
		return ans
	}

	return traverse(root, root.Val)
}
