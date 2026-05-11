/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, data *[]int, level int) {
	if root == nil {
		return
	}
	if len(*data) < level + 1 {
		*data = append(*data, root.Val)
	}
	core(root.Right, data, level + 1)
	core(root.Left, data, level + 1)
}

func rightSideView(root *TreeNode) []int {
    ans := []int{}
	core(root, &ans, 0)
	return ans
}
