package pascals_triangle_ii

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := getRow(rowIndex - 1)

	newRow := make([]int, rowIndex+1)
	newRow[0], newRow[rowIndex] = 1, 1

	for i := 1; i < rowIndex; i++ {
		newRow[i] = prevRow[i-1] + prevRow[i]
	}

	return newRow
}
