/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var ans *ListNode

	runner := head
	for runner != nil {
		tmp := runner.Next

		runner.Next = ans
		ans = runner

		runner = tmp
	}

	return ans
}
