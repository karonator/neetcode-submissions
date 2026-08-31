/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}
	slow := dummy
	fast := dummy
	counter := 0

	for fast.Next != nil {
		fast = fast.Next
		counter ++
		if counter > n {
			slow = slow.Next
		}
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
