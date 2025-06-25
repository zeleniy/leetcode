package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer bool
	nums   []int
}

func BenchmarkContainsDuplicate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			containsDuplicate(data.nums)
		}
	}
}

func TestContainsDuplicate(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, containsDuplicate(data.nums))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{true, []int{1, 2, 3, 1}},
		{false, []int{1, 2, 3, 4}},
		{true, []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
		{false, []int{1, 5, -2, -4, 0}},
		{true, []int{2, 14, 18, 22, 22}},
	}
}
