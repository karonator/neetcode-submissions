/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		return core(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return core(root.Left, p, q)
	}
	return root
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	return core(root, p, q)
}
