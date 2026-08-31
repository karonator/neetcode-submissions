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

	var ans *ListNode
	var tail *ListNode

	for runner1 != nil && runner2 != nil {
		val := runner1.Val + runner2.Val
		if carry {
			val++
		}
		
		if val > 9 {
			val = val % 10
			carry = true
		} else {
			carry = false
		}

		if ans == nil {
			ans = &ListNode{
				Val: val,
			}
			tail = ans
		} else {
			tail.Next = &ListNode{
				Val: val,
			}
			tail = tail.Next
		}

		runner1 = runner1.Next
		runner2 = runner2.Next
	}

	var left *ListNode
	if runner1 != nil {
		left = runner1
	}
	if runner2 != nil {
		left = runner2
	}
	for left != nil {
		val := left.Val
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
		left = left.Next
	}

	if carry {
		tail.Next = &ListNode{
			Val: 1,
			Next: nil,
		}
	}

	return ans
}
