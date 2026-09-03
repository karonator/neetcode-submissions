/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, ans *[]int, level int) {
	if root == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, root.Val)
	}
	dfs(root.Right, ans, level + 1)
	dfs(root.Left, ans, level + 1)
}

func rightSideView(root *TreeNode) []int {
    ans := make([]int, 0)
	dfs(root, &ans, 0)
	return ans
}
