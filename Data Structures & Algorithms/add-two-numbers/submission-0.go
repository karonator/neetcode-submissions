/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    runner_1 := l1
    runner_2 := l2
	carry := 0

	var head *ListNode
	var last *ListNode

	for runner_1 != nil || runner_2 != nil || carry != 0 {
		val := carry
		if runner_1 != nil {
			val += runner_1.Val
			runner_1 = runner_1.Next
		}
		if runner_2 != nil {
			val += runner_2.Val
			runner_2 = runner_2.Next
		}
		if val > 9 {
			val = val % 10
			carry = 1
		} else {
			carry = 0
		}
		newbie := &ListNode{
			Val: val,
		}
		if head == nil {
			head = newbie
			last = head
		} else {
			last.Next = newbie
			last = last.Next
		}
	}
	return head
}
