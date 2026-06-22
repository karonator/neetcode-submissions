import (
    "slices"
)

func threeSum(nums []int) [][]int {
	const MINUS_BIG_NUM = -10000000

	slices.Sort(nums)

	last_i_val := MINUS_BIG_NUM
	result := [][]int{}

	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] == last_i_val {
			continue
		}
		last_i_val = nums[i]

		j := i + 1
		k := len(nums) - 1

		last_j_val := MINUS_BIG_NUM
		last_k_val := MINUS_BIG_NUM

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				if nums[j] == last_j_val {
					j++
					continue
				}
				if nums[k] == last_k_val {
					k--
					continue
				}
				result = append(result, []int{nums[i], nums[j], nums[k]})
				last_j_val = nums[j]
				last_k_val = nums[k]
				j++
			}
		}
	}

	return result
}
