/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	core(root.Left, ans)
	*ans = append(*ans, root.Val)
	core(root.Right, ans)
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	core(root, &ans)

	return ans
}
