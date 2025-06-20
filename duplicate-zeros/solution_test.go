package duplicate_zeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	arr    []int
}

func BenchmarkDuplicateZeros(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			duplicateZeros(data.arr)
		}
	}
}

func TestDuplicateZeros(t *testing.T) {
	for _, data := range getTestDataSet() {
		duplicateZeros(data.arr)
		assert.Equal(t, data.answer, data.arr)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 0, 0}, []int{1, 0, 2}},
		{[]int{1, 0, 0, 2, 3, 0, 0, 4}, []int{1, 0, 2, 3, 0, 4, 5, 0}},
		{[]int{1, 0, 0, 2, 3, 0, 0, 4}, []int{1, 0, 2, 3, 0, 4, 5, 0}},
		{[]int{8, 4, 5, 0, 0, 0, 0, 0}, []int{8, 4, 5, 0, 0, 0, 0, 7}},
		{[]int{0, 0, 1, 7, 6, 0, 0, 2}, []int{0, 1, 7, 6, 0, 2, 0, 7}},
		{[]int{0, 0, 8, 0}, []int{0, 8, 0, 0}},
	}
}
