package minimum_operations_to_reduce_x_to_zero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	x      int
	nums   []int
}

func BenchmarkMinOperations(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			minOperations(data.nums, data.x)
		}
	}
}

func TestMinOperations(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, minOperations(data.nums, data.x))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{2, 5, []int{1, 1, 4, 2, 3}},
		{-1, 4, []int{5, 6, 7, 8, 9}},
		{5, 10, []int{3, 2, 20, 1, 1, 3}},
	}
}
