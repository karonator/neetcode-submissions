/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func core(root *TreeNode, cnt *int, k int) (int, bool) {
	if root == nil {
		return 0, false
	}

	left, lfound := core(root.Left, cnt, k)
	if lfound {
		return left, true
	}

	*cnt ++
	if *cnt == k {
		return root.Val, true
	}

	right, rfound := core(root.Right, cnt, k)
	if rfound {
		return right, true
	}
	return 0, false
}

func kthSmallest(root *TreeNode, k int) int {
	cnt := 0
	ans, _ := core(root, &cnt, k)
	return ans
}
