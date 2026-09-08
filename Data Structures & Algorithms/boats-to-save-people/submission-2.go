func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i int, j int) bool {
		return people[i] < people[j]
	})

	i := 0
	j := len(people) - 1
	boats := 0

	for i <= j {
		space := limit - people[j]
		j--
		if space >= people[i]	{
			i++
		}
		boats++
	}
	return boats
}
