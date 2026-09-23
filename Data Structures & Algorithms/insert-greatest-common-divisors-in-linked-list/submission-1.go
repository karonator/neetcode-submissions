/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    gcd := func(a int, b int) int {
		x := max(a, b)
		y := min(a, b)

		for y != 0 {
			x, y = y, x % y
		}
		return x
	}

	runner := head
	for runner != nil {
		if runner.Next != nil {
			node := &ListNode {
				Val: gcd(runner.Val, runner.Next.Val),
				Next: runner.Next,
			}
			runner.Next = node
			runner = runner.Next
		}
		runner = runner.Next
	}
	return head
}
