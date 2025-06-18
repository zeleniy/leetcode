package keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {

	n := len(rooms)

	visited := make([]bool, n)
	visited[0] = true

	stack := make([]int, 1, 4)
	stack[0] = 0

	for len(stack) > 0 {

		key := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, key := range rooms[key] {
			if !visited[key] {
				stack = append(stack, key)
				visited[key] = true
				n--
			}
		}
	}

	return n == 1
}
