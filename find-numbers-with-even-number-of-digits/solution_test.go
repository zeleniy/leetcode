package find_numbers_with_even_number_of_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	nums   []int
}

func TestFindNumbers(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, findNumbers(dataSet.nums))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{2, []int{12, 345, 2, 6, 7896}},
		{1, []int{555, 901, 482, 1771}},
	}
}
