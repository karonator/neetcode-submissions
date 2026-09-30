import (
	"slices"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func dist(p1 []int, p2 []int) int {
	return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])
}

type Edge struct {
	Dist int
	Start int
	End int
}

func minCostConnectPoints(points [][]int) int {
	edges := []Edge{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				Start: i,
				End: j,
				Dist: dist(points[i], points[j]),
			})
		}
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		return a.Dist - b.Dist
	})

	parents := make([]int, len(points))
	for i := range parents {
		parents[i] = i
	}

	find := func(a int) int {
		path := []int{}
		for parents[a] != a {
			path = append(path, a)
			a = parents[a]
		}
		for _, i := range path {
			parents[i] = a
		}
		return a		
	}

	union := func(a, b int) bool {
		ra, rb := find(a), find(b)
		if ra != rb {
			parents[rb] = ra
			return true
		} else {
			return false
		}
	}

	ans := 0
	for _, edge := range edges {
		if union(edge.Start, edge.End) {
			ans += edge.Dist
		}
	}

	return ans
}
