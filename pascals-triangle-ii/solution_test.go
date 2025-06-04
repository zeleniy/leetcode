package pascals_triangle_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	numRows int
	answer  []int
}

func TestGetRow(t *testing.T) {

	for _, dataSet := range getMergeData() {

		triangle := getRow(dataSet.numRows)

		assert.Equal(t, len(dataSet.answer), len(triangle))
		assert.Equal(t, dataSet.answer, triangle)
	}
}

func getMergeData() []TestCaseDataSet {

	return []TestCaseDataSet{
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
	}
}
