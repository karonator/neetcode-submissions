/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func core(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }

    left_depth, left_ok := core(root.Left)
    right_depth, right_ok := core(root.Right)

    if !left_ok || !right_ok {
        return 0, false
    }

    if abs(left_depth - right_depth) > 1 {
        return 0, false
    }

    return max(left_depth, right_depth) + 1, true
}


func isBalanced(root *TreeNode) bool {
    _, ans := core(root)
    return ans
}
