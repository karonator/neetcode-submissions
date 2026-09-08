func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i int, j int) bool {
		return people[i] < people[j]
	})

	i := 0
	j := len(people) - 1
	boats := 0

	for i <= j {
		cap := limit - people[j]
		j--
		if cap >= people[i]	{
			i++
		}
		boats++
	}
	return boats
	// 4
	// 1 1 2 3 4
}
