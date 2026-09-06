/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
   	var del func(*TreeNode) *TreeNode
	del = func(node *TreeNode) (newRoot *TreeNode) {
		if node == nil {
			return nil
		}
		if (node.Val == key) {
			if node.Left == nil && node.Right == nil {
				return nil
			}
			if node.Left != nil && node.Right != nil {
				runner := node.Right
				for runner.Left != nil {
					runner = runner.Left
				}
				runner.Left = node.Left
				return node.Right
			}
			if node.Right != nil {
				return node.Right
			}
			return node.Left
		} else {
			node.Left = del(node.Left)
			node.Right = del(node.Right)
			return node
		}
	}
	return del(root)
}
