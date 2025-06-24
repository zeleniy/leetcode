package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	target int
	nums   []int
}

func BenchmarkTwoSum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			twoSum(data.nums, data.target)
		}
	}
}

func TestTwoSum(t *testing.T) {
	for _, data := range getTestDataSet() {
		assert.ElementsMatch(t, data.answer, twoSum(data.nums, data.target))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{0, 1}, 9, []int{2, 7, 11, 15}},
		{[]int{1, 2}, 6, []int{3, 2, 4}},
		{[]int{0, 1}, 6, []int{3, 3}},
	}
}
