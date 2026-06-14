func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// finding center
	fast := head
	slow := head
	var superslow *ListNode

	for fast != nil && fast.Next != nil {
		if superslow == nil {
			superslow = slow
		} else {
			superslow = superslow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	superslow.Next = nil

	// reverse second half
	var reverted *ListNode
	runner := slow

	for runner != nil {
		tmp := runner.Next
		runner.Next = reverted
		reverted = runner
		runner = tmp
	}

	// merge
	runner1 := head
	runner2 := reverted

	for runner1 != nil && runner2 != nil {
		next1 := runner1.Next
		next2 := runner2.Next

		runner1.Next = runner2

		if next1 == nil {
			break
		}

		runner2.Next = next1

		runner1 = next1
		runner2 = next2
	}
}