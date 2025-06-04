package pascals_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	numRows int
	answer  [][]int
}

func TestGenerate(t *testing.T) {

	for _, dataSet := range getMergeData() {

		triangle := generate(dataSet.numRows)

		assert.Equal(t, len(dataSet.answer), len(triangle))
		assert.Equal(t, dataSet.answer, triangle)
	}
}

func getMergeData() []TestCaseDataSet {

	return []TestCaseDataSet{
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
}
