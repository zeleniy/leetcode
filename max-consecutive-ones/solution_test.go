package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer   int
	sequence []int
}

func TestFib(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, findMaxConsecutiveOnes(dataSet.sequence))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{0, []int{0}},
		{1, []int{1}},
		{3, []int{1, 1, 0, 1, 1, 1}},
		{2, []int{1, 0, 1, 1, 0, 1}},
	}
}
