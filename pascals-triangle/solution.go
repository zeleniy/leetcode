package pascals_triangle

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	}

	prevRows := generate(numRows - 1)
	lastRow := prevRows[len(prevRows)-1]

	newRow := make([]int, numRows)
	newRow[0], newRow[numRows-1] = 1, 1

	for i := 1; i < numRows-1; i++ {
		newRow[i] = lastRow[i-1] + lastRow[i]
	}

	return append(prevRows, newRow)
}
