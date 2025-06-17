package zero_one_matrix

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func updateMatrix(mat [][]int) [][]int {

	m, n := len(mat), len(mat[0])
	result := make([][]int, m)
	maxDistance := m + n

	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = maxDistance
		}
	}

	queue := make([][2]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			x, y := cell[0]+dir[0], cell[1]+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if result[x][y] > result[cell[0]][cell[1]]+1 {
					result[x][y] = result[cell[0]][cell[1]] + 1
					queue = append(queue, [2]int{x, y})
				}
			}
		}
	}

	return result
}
