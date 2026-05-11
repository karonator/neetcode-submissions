/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func builder(preorder []int, preorder_index *int, inorder_map map[int]int, l_bound int, r_bound int) *TreeNode {
    if l_bound == r_bound {
        return nil
    }
    node := &TreeNode { Val: preorder[*preorder_index] }
    *preorder_index++

    node.Left = builder(preorder, preorder_index, inorder_map, l_bound, inorder_map[node.Val])
    node.Right = builder(preorder, preorder_index, inorder_map, inorder_map[node.Val] + 1, r_bound)

    return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    inorder_map := make(map[int]int)
    for i, v := range(inorder) {
        inorder_map[v] = i
    }
    preorder_index := 0
    return builder(preorder, &preorder_index, inorder_map, 0, len(preorder))
}
