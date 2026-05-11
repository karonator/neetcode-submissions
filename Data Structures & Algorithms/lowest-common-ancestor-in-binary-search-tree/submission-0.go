/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func search(node *TreeNode, min int, max int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val <= max && node.Val >= min {
		return node
	}
	left := search(node.Left, min, max)
	if left != nil {
		return left
	}
	right := search(node.Right, min, max)
	if right != nil {
		return right
	}
	return nil
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    min := q.Val
	max := p.Val
	if p.Val < q.Val {
		min = p.Val
		max = q.Val
	}
	return search(root, min, max)
}
