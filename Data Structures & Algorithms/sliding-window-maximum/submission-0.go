func maxSlidingWindow(nums []int, k int) []int {
    deque := []int{}
	ans := []int{}
	for right := 0; right < len(nums); right++ {
		// 1. Удалить из front всё, что вышло из окна
		for len(deque) > 0 {
			if deque[0] < right - k + 1 {
				deque = deque[1:]
			} else {
				break
			}
		}
    	
		// 2. Почистить back:
    	//    пока nums[deque[last]] <= nums[right]
		for len(deque) > 0 {
			if nums[deque[len(deque) - 1]] <= nums[right] {
				deque = deque[:len(deque) - 1]
			} else {
				break
			}
		}

		deque = append(deque, right)

		if right >= k - 1 {
			ans = append(ans, nums[deque[0]])
		}
	}

	return ans
}
