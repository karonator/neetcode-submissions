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
	carry := false

	dummy := &ListNode{}
	tail := dummy

	for runner1 != nil || runner2 != nil {
		val := 0
		if runner1 != nil {
			val += runner1.Val
			runner1 = runner1.Next
		}
		if runner2 != nil {
			val += runner2.Val
			runner2 = runner2.Next
		}
		if carry {
			val++
		}
		
		if val > 9 {
			val = val % 10
			carry = true
		} else {
			carry = false
		}

		tail.Next = &ListNode{
			Val: val,
		}
		tail = tail.Next
	}

	if carry {
		tail.Next = &ListNode{
			Val: 1,
			Next: nil,
		}
	}

	return dummy.Next
}
