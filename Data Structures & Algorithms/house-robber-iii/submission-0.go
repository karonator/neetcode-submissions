/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (rootRobbed int, rootNotRobbed int) {
		if node == nil {
			return 0, 0
		}
		lR, lN := dfs(node.Left)
		rR, rN := dfs(node.Right)
		return node.Val + lN + rN, max(lR, lN) + max(rR, rN)
	}

	ans1, ans2 := dfs(root)
	return max(ans1, ans2)
}
