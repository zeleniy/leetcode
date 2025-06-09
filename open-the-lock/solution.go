package open_the_lock

func openLock(deadends []string, target string) int {

	visited := make(map[string]bool)

	for _, deadend := range deadends {
		visited[deadend] = true
	}

	if visited["0000"] {
		return -1
	}

	queue := []string{"0000"}
	visited["0000"] = true

	for moves := 0; len(queue) > 0; moves++ {
		for levelSize := len(queue); 0 < levelSize; levelSize-- {

			current := queue[0]
			queue = queue[1:]

			if current == target {
				return moves
			}

			for _, newNumber := range getNeighbors(current) {
				if _, exists := visited[newNumber]; !exists {
					visited[newNumber] = true
					queue = append(queue, newNumber)
				}
			}
		}
	}

	return -1
}

func getNeighbors(number string) [8]string {

	neighbors := [8]string{}

	for i := 0; i < 4; i++ {

		digit := int(number[i] - '0')

		neighbors[i*2] = modifyDigit(number, digit, i, 1)
		neighbors[i*2+1] = modifyDigit(number, digit, i, -1)
	}

	return neighbors
}

func modifyDigit(number string, digit, i, delta int) string {

	bytes := []byte(number)
	bytes[i] = byte((digit+delta+10)%10 + '0')

	return string(bytes)
}
