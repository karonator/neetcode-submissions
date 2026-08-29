/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
	var tail *ListNode

	runner1 := list1
	runner2 := list2

	for runner1 != nil || runner2 != nil {
		if runner1 == nil || runner2 != nil && runner1 != nil && runner2.Val <= runner1.Val {
			if head == nil {
				head = runner2
				tail = head
			} else {
				tail.Next = runner2
				tail = tail.Next
			}
			runner2 = runner2.Next
		} else {
			if head == nil {
				head = runner1
				tail = head
			} else {
				tail.Next = runner1
				tail = tail.Next
			}
			runner1 = runner1.Next
		}
	}
	return head
}
