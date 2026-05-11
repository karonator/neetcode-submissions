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
	if len(*data) < level + 1 {
		*data = append(*data, []int{})
	}
	(*data)[level] = append((*data)[level], root.Val)
	core(root.Left, data, level + 1)
	core(root.Right, data, level + 1)
}

func levelOrder(root *TreeNode) [][]int {
    answer := [][]int{}
	core(root, &answer, 0)
	return answer
}
