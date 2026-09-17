/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow := head
	fast := head

	if head == nil {
		return
	}

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	center := slow.Next
	slow.Next = nil

	var reversed *ListNode
	runner := center
	for runner != nil {
		tmp := runner.Next
		runner.Next = reversed
		reversed = runner
		runner = tmp
	}

	first := head
	second := reversed
	for second != nil {
		n1 := first.Next
		n2 := second.Next

		first.Next = second
		second.Next = n1

		first = n1
		second = n2
	}
}
