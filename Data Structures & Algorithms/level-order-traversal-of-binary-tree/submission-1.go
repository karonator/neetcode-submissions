/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, data *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*data) <= level {
		*data = append(*data, make([]int, 0))
	}
	(*data)[level] = append((*data)[level], root.Val)

	core(root.Left, data, level + 1)
	core(root.Right, data, level + 1)
}

func levelOrder(root *TreeNode) [][]int {
    ans := make([][]int, 0)
	core(root, &ans, 0)
	return ans
}
