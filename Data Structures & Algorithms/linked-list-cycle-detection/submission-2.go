/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}
