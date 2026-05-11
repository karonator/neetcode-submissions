/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func find_left(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return find_left(node.Left)
	}
	return node
}

func delete(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val != key {
		if key > node.Val {
			node.Right = delete(node.Right, key)
		} else {
			node.Left = delete(node.Left, key)
		}
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		if node.Left != nil && node.Right != nil {
			new_root := node.Right
			most_left := find_left(new_root)
			most_left.Left = node.Left
			return new_root
		} else {
			if node.Left == nil {
				return node.Right
			} else {
				return node.Left
			}
		}
	}
	return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    return delete(root, key)
}
