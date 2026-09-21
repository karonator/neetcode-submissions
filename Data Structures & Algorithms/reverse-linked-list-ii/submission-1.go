/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{
		Val: -1,
		Next: head,
	}
	var stash *ListNode
	runner := dummy

	for i := 0; i < left - 1; i++ {
		runner = runner.Next
	}
	stash = runner

	runner = runner.Next
	var xHead *ListNode
	xTail := runner

	for i := 0; i < right - left + 1; i++ {
		tmp := runner.Next
		runner.Next = xHead
		xHead = runner
		runner = tmp
	}

	stash.Next = xHead
	xTail.Next = runner

	return dummy.Next
}
