/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    runner1 := l1
	runner2 := l2
	carry := 0

	dummy := &ListNode{}
	tail := dummy

	for runner1 != nil || runner2 != nil || carry > 0 {
		val := carry
		if runner1 != nil {
			val += runner1.Val
			runner1 = runner1.Next
		}
		if runner2 != nil {
			val += runner2.Val
			runner2 = runner2.Next
		}

		carry = val / 10
		val = val % 10

		tail.Next = &ListNode{
			Val: val,
		}
		tail = tail.Next
	}

	return dummy.Next
}
