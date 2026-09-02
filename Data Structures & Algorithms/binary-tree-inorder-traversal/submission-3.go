/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	type Item struct {
		Node 	*TreeNode
		Visited bool
	}

	ans := make([]int, 0)

	stack := make([]Item, 0)
	stack = append(stack, Item{
		Node: root,
		Visited: false,
	})

	for len(stack) > 0 {
		elem := stack[len(stack) - 1]
		stack = stack[:len(stack) -1]

		if elem.Node == nil {
			continue
		}

		if elem.Visited {
			ans = append(ans, elem.Node.Val)
		} else {
			stack = append(stack, Item{
				Node: elem.Node.Right,
				Visited: false,
			})
			stack = append(stack, Item{
				Node: elem.Node,
				Visited: true,
			})
			stack = append(stack, Item{
				Node: elem.Node.Left,
				Visited: false,
			})
		}
	}

	return ans
}
