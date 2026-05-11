/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(node *TreeNode, data *[]int) {
	if node == nil {
		return
	}

	core(node.Left, data)
	*data = append(*data, node.Val)
	core(node.Right, data)
}

func inorderTraversal(root *TreeNode) []int {
	answer := []int{}
	core(root, &answer)
	return answer
}
