/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func isBalanced(root *TreeNode) bool {
	var check func(*TreeNode) (int, bool)
	check = func(node *TreeNode) (depth int, balanced bool) {
		if node == nil {
			return 0, true
		}
		lD, lOk := check(node.Left)
		rD, rOk := check(node.Right)

		return max(lD, rD) + 1, lOk && rOk && abs(lD - rD) < 2
	}
	_, ok := check(root)
	return ok
}
