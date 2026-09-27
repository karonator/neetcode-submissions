func findOrder(numCourses int, prerequisites [][]int) []int {
    indegree := make(map[int]map[int]struct{})
    deps := make(map[int]int)

    for i := range prerequisites {
        if _, found := indegree[prerequisites[i][1]]; !found {
            indegree[prerequisites[i][1]] = make(map[int]struct{})
        }
        indegree[prerequisites[i][1]][prerequisites[i][0]] = struct{}{}
        deps[prerequisites[i][0]]++
    }

    queue := make([]int, 0)
    for i := range numCourses {
        if deps[i] == 0 {
            queue = append(queue, i)
        }
    }

	ans := []int{}
    deleted := 0
    for len(queue) > 0 {
        cur := queue[0]
		ans = append(ans, cur)
        queue = queue[1:]
        deleted++

        for i := range indegree[cur] {
            deps[i]--
            if deps[i] == 0 {
                queue = append(queue, i)
            }
        }
    }

    if deleted < numCourses {
        return []int{}
    }
    return ans
}
