func flood(grid [][]int, x int, y int) int {
    if grid[x][y] == 2 || grid[x][y] != 1 {
        return 0
    }
	ans := 1
    grid[x][y] = 2

    if x - 1 >= 0 {
        ans += flood(grid, x - 1, y)
    }
    if y - 1 >= 0 {
        ans += flood(grid, x, y - 1)
    }
    if x + 1 < len(grid) {
        ans += flood(grid, x + 1, y)
    }
    if y + 1 < len(grid[0]) {
        ans += flood(grid, x, y + 1)
    }
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grid [][]int) int {
    max_area := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                continue
            }
            if grid[i][j] == 1 {
                max_area = max(max_area, flood(grid, i, j))
            }
        }   
    }

    return max_area    
}
