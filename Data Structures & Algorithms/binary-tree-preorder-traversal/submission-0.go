/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

	*data = append(*data, node.Val)

	core(node.Left, data)
	core(node.Right, data)
}

func preorderTraversal(root *TreeNode) []int {
	answer := []int{}
	core(root, &answer)
	return answer    
}
