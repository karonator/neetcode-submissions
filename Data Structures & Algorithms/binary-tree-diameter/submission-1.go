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

func core(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    left_depth, left_dia := core(root.Left)
    right_depth, right_dia := core(root.Right)
    depth := max(left_depth, right_depth) + 1
    diam := max(max(left_depth + right_depth, left_dia), right_dia)

    return depth, diam
}

func diameterOfBinaryTree(root *TreeNode) int {
    _, diam := core(root)

    return diam
}
