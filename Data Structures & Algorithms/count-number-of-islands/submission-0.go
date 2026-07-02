func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
		return 0
	}
	m,n := len(grid), len(grid[0])
	isLands := make([][]bool,m)
	for i,_ := range isLands {
		isLands[i] = make([]bool,n)
	}
	count := 0
	for i := 0; i <len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && isLands[i][j] == false {
				count += 1
				queue := [][]int {{i,j}}
				for len(queue) > 0 {
					x,y := queue[0][0], queue[0][1]
					isLands[x][y] = true
					queue = queue[1:]
					if (x + 1) < len(grid) && isLands[x+1][y] == false && grid[x+1][y] == '1' {
						queue = append(queue,[]int{x+1,y})
					}
					if (y + 1) < len(grid[x]) && isLands[x][y+1] == false && grid[x][y+1] == '1' {
						queue = append(queue,[]int{x,y+1})
					}
					if (x-1) >= 0 && isLands[x-1][y] == false && grid[x-1][y] == '1' {
						queue = append(queue,[]int{x-1,y})
					}
					if (y-1) >= 0 && isLands[x][y-1] == false && grid[x][y-1] == '1' {
						queue = append(queue,[]int{x,y-1})
					}
				}
			} else {
				isLands[i][j] = true
			}
		}
	}
	return count
}
