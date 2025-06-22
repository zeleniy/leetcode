package open_the_lock

var neighbors [8]int
var forward = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var backward = [10]int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}

func openLock(deadends []string, target string) int {

	targetInt := toInt(target)
	var visited [10000]bool

	for _, deadend := range deadends {
		visited[toInt(deadend)] = true
	}

	if visited[0] {
		return -1
	}

	queue := make([]int, 1, 8192)
	queue[0] = 0
	visited[0] = true

	for moves := 0; len(queue) > 0; moves++ {
		for levelSize := len(queue); 0 < levelSize; levelSize-- {

			current := queue[0]
			queue = queue[1:]

			if current == targetInt {
				return moves
			}

			populateCombinations(current)
			for i := 0; i < 8; i++ {
				if newNumber := neighbors[i]; !visited[newNumber] {
					visited[newNumber] = true
					queue = append(queue, newNumber)
				}
			}
		}
	}

	return -1
}

func populateCombinations(num int) {

	d0 := num % 10
	d1 := (num / 10) % 10
	d2 := (num / 100) % 10
	d3 := num / 1000

	neighbors[0] = num - d0 + forward[d0]
	neighbors[1] = num - d0 + backward[d0]
	neighbors[2] = num - d1*10 + forward[d1]*10
	neighbors[3] = num - d1*10 + backward[d1]*10
	neighbors[4] = num - d2*100 + forward[d2]*100
	neighbors[5] = num - d2*100 + backward[d2]*100
	neighbors[6] = num - d3*1000 + forward[d3]*1000
	neighbors[7] = num - d3*1000 + backward[d3]*1000
}

func toInt(s string) int {

	return int(s[0]-'0')*1000 + int(s[1]-'0')*100 + int(s[2]-'0')*10 + int(s[3]-'0')
}
