package flood_fill

var neighbors = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	m := len(image)
	n := len(image[0])
	originalColor := image[sr][sc]

	if originalColor == color {
		return image
	}

	queue := [][2]int{{sr, sc}}
	image[sr][sc] = color

	for len(queue) > 0 {

		i, j := queue[0][0], queue[0][1]

		for _, neighbor := range neighbors {
			y, x := i-neighbor[0], j-neighbor[1]
			if y >= 0 && y < m && x >= 0 && x < n && image[y][x] == originalColor {
				image[y][x] = color
				queue = append(queue, [2]int{y, x})
			}
		}

		queue = queue[1:]
	}

	return image
}
