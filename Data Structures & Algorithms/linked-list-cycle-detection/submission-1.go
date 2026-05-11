/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

    slow := head.Next
	fast := head.Next
	if fast == nil {
		return false
	}
	fast = fast.Next

	for fast != nil {
		if slow.Val == fast.Val {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
	}

	return false
}
