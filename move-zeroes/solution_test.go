package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	nums   []int
}

func BenchmarkMoveZeroes(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			moveZeroes(data.nums)
		}
	}
}

func TestMoveZeroes(t *testing.T) {

	for _, data := range getTestDataSet() {
		moveZeroes(data.nums)
		assert.Equal(t, data.answer, data.nums)
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 3, 12, 0, 0}, []int{0, 1, 0, 3, 12}},
		{[]int{0}, []int{0}},
		{[]int{1, 0, 0}, []int{0, 0, 1}},
	}
}
