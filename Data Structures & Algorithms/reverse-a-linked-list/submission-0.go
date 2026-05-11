/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	runner := head

	for runner != nil {
		new := &ListNode {
			Val : runner.Val,
		}
		if result != nil {
			new.Next = result
		}
		result = new
		runner = runner.Next
	}

	return result
}
