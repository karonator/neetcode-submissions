/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    first := head
	var second *ListNode
	cnt := 0

	for first != nil {
		fmt.Println(first)
		if cnt == n {
			second = head
		}
		if cnt > n {
			second = second.Next
		}
		cnt++
		first = first.Next
	}

	if second == nil {
		return head.Next
	}
	second.Next = second.Next.Next

	return head
}
