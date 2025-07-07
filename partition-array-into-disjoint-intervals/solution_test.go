package partition_array_into_disjoint_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	nums   []int
}

func BenchmarkPartitionDisjoint(b *testing.B) {

	dataSet := getTestDataSet()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			partitionDisjoint(data.nums)
		}
	}
}

func TestPartitionDisjoint(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, partitionDisjoint(data.nums))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{3, []int{5, 0, 3, 8, 6}},
		{4, []int{1, 1, 1, 0, 6, 12}},
		{9, []int{90, 47, 69, 10, 43, 92, 31, 73, 61, 97}},
		{7, []int{32, 57, 24, 19, 0, 24, 49, 67, 87, 87}},
	}
}
