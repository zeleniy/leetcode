package squares_of_a_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	nums   []int
}

func BenchmarkSortedSquares(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			sortedSquares(data.nums)
		}
	}
}

func TestSortedSquares(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, sortedSquares(data.nums))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 4, 9}, []int{-2, 1, 3}},
		{[]int{1, 4, 9}, []int{-3, -2, -1}},
		{[]int{0, 1, 9, 16, 100}, []int{-4, -1, 0, 3, 10}},
		{[]int{4, 9, 9, 49, 121}, []int{-7, -3, 2, 3, 11}},
	}
}
