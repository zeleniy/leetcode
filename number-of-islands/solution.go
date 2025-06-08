package reverse_string

var neighbors = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	counter := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				grid[i][j] = '0'
				markIsland(grid, m, n, i, j)
			}
		}
	}

	return counter
}

func markIsland(grid [][]byte, m, n, i, j int) {

	queue := [][]int{{i, j}}

	for len(queue) > 0 {

		i, j := queue[0][0], queue[0][1]

		for _, neighbor := range neighbors {
			y, x := i-neighbor[0], j-neighbor[1]
			if y >= 0 && y < m && x >= 0 && x < n {
				if grid[y][x] == '1' {
					grid[y][x] = '0'
					queue = append(queue, []int{y, x})
				}
			}
		}

		queue = queue[1:]
	}
}
