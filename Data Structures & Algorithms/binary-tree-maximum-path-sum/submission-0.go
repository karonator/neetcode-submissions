/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    ans := math.MinInt

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		pathViaNode := node.Val + max(left, 0) + max(right, 0)
		ans = max(ans, pathViaNode)
		return node.Val + max(left, right, 0)
	}
	dfs(root)
	return ans
}
