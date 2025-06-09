package perfect_squares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	n      int
}

func TestNumSquares(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, numSquares(dataSet.n))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{1, 1},
		{1, 4},
		{4, 7},
		{3, 12},
		{2, 13},
		{4, 23},
		{1, 100},
		{2, 500},
		{2, 1000},
		{2, 5000},
		{1, 10000},
		{3, 12345},
		{4, 9999},
	}
}
