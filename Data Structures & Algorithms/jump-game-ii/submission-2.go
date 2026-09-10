func jump(nums []int) int {
	jumps := 0
	// граница текущего прыжка
	end := 0
	// максимально достижимая граница
	far := 0

	for i := 0; i < len(nums) - 1; i++ {
		far = max(far, i + nums[i])
		if i == end {
			jumps++
			end = far
		}
	}
	return jumps
}
