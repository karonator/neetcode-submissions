/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    runner1 := list1
	runner2 := list2

	var result *ListNode
	var last * ListNode

	for runner1 != nil && runner2 != nil {
		if runner1.Val < runner2.Val {
			if result == nil {
				result = runner1
				last = result
			} else {
				last.Next = runner1
				last = last.Next
			}
			runner1 = runner1.Next
		} else {
			if result == nil {
				result = runner2
				last = result
			} else {
				last.Next = runner2
				last = last.Next
			}
			runner2 = runner2.Next
		}
	}

	if runner1 == nil {
		if result == nil {
			result = runner2
		} else {
			last.Next = runner2
		}
	}

	if runner2 == nil {
		if result == nil {
			result = runner1
		} else {
			last.Next = runner1
		}
	}

	return result
}
