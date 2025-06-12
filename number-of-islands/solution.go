package number_of_islands

var neighbors = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numIslandsDfs(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	counter := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				markIslandDfs(grid, m, n, i, j)
			}
		}
	}

	return counter
}

func markIslandDfs(grid [][]byte, m, n, i, j int) {

	stack := make([][2]int, 1, 8)
	stack[0] = [2]int{i, j}
	grid[i][j] = '0'

	for len(stack) > 0 {

		stackLength := len(stack)

		i, j := stack[stackLength-1][0], stack[stackLength-1][1]
		stack = stack[:stackLength-1]

		for _, neighbor := range neighbors {
			y, x := i-neighbor[0], j-neighbor[1]
			if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] == '1' {
				grid[y][x] = '0'
				stack = append(stack, [2]int{y, x})
			}
		}
	}
}

func numIslandsBfs(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	counter := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				markIslandBfs(grid, m, n, i, j)
			}
		}
	}

	return counter
}

func markIslandBfs(grid [][]byte, m, n, i, j int) {

	queue := make([][2]int, 1, 8)
	queue[0] = [2]int{i, j}
	grid[i][j] = '0'

	for len(queue) > 0 {

		i, j := queue[0][0], queue[0][1]

		for _, neighbor := range neighbors {
			y, x := i-neighbor[0], j-neighbor[1]
			if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] == '1' {
				grid[y][x] = '0'
				queue = append(queue, [2]int{y, x})
			}
		}

		queue = queue[1:]
	}
}
