func longestConsecutive(nums []int) int {
	lib := make(map[int]struct{})
	for _, n := range nums {
		lib[n] = struct{}{}
	}

	heads := make(map[int]struct{})
	for n, _ := range lib {
		if _, found := lib[n - 1]; !found {
			heads[n] = struct{}{}
		}
	}

	result := 0
	for head, _ := range heads {
		seq_len := 1
		runner := head
		for {
			_, found := lib[runner + 1]
			if !found {
				break
			}
			seq_len++
			runner++
		}
		if result < seq_len {
			result = seq_len
		}
	}

	return result
}
