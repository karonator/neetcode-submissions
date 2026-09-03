/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, k int, cur *int, ans *int) {
	if root == nil || *cur > k {
		return
	}
	core(root.Left, k, cur, ans)
	if k == *cur {
		*ans = root.Val
	}
	*cur++
	core(root.Right, k, cur, ans)
}

func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	cur := 1
	core(root, k, &cur, &ans)
	return ans
}
