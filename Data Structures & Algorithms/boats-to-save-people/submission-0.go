import (
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	ans := 0
	i := 0
	j := len(people) - 1

	for i <= j {
		free := limit - people[j]
		j--
		if free >= people[i] {
			i++
		}
		ans++
	}
	return ans
}
