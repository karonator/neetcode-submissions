func flood(grid [][]byte, x int, y int) {
    if grid[x][y] == '2' || grid[x][y] != '1' {
        return
    }
    grid[x][y] = '2'
    if x - 1 >= 0 {
        flood(grid, x - 1, y)
    }
    if y - 1 >= 0 {
        flood(grid, x, y - 1)
    }
    if x + 1 < len(grid) {
        flood(grid, x + 1, y)
    }
    if y + 1 < len(grid[0]) {
        flood(grid, x, y + 1)
    }
}

func numIslands(grid [][]byte) int {
    islands := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '2' {
                continue
            }
            if grid[i][j] == '1' {
                islands++
                flood(grid, i, j)
            }
            
        }   
    }

    return islands
}
