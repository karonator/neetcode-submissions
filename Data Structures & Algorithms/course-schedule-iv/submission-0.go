func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	reach := make([][]bool, numCourses)
	for i := range reach {
		reach[i] = make([]bool, numCourses)
	}

	for _, prerequisite := range prerequisites {
		from := prerequisite[0]
		to := prerequisite[1]
		reach[from][to] = true
	}

	for i := range numCourses {
		for j := range numCourses {
			for k := range numCourses {
				reach[j][k] = reach[j][k] || (reach[j][i] && reach[i][k])
			}
		}
	}

	answer := make([]bool, len(queries))
	for i, query := range queries {
		answer[i] = reach[query[0]][query[1]]
	}

	return answer
}
