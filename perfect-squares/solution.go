package perfect_squares

func numSquares(n int) int {

	queue := []int{n}
	visited := make([]bool, n+1)
	precedingSquares := getNeighbors(n)

	for steps := 0; len(queue) > 0; steps++ {
		for levelSize := len(queue); 0 < levelSize; levelSize-- {

			current := queue[0]
			queue = queue[1:]

			for _, neighbor := range precedingSquares {

				difference := current - neighbor

				if difference == 0 {
					return steps + 1
				}

				if difference > 0 && !visited[difference] {
					visited[difference] = true
					queue = append(queue, difference)
				}
			}
		}
	}

	return -1
}

func getNeighbors(n int) []int {

	neighbors := []int{}

	for i := 1; ; i++ {
		if square := i * i; square <= n {
			neighbors = append(neighbors, square)
		} else {
			return neighbors
		}
	}
}
