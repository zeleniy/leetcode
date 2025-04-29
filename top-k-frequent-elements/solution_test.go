package top_k_frequent_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	expected []int
	nums     []int
	k        int
}

func TestMerge(t *testing.T) {

	for _, dataSet := range getMergeData() {
		assert.Equal(t, dataSet.expected, topKFrequent(dataSet.nums, dataSet.k))
	}
}

func getMergeData() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1}, []int{1, 1, 1, 2, 2, 3}, 1},
		{[]int{1, 2}, []int{3, 2, 2, 1, 1, 1}, 2},
		{[]int{1, 2}, []int{1, 1, 1, 2, 2, 3}, 2},
		{[]int{1, 2, 3}, []int{1, 1, 1, 2, 2, 3}, 3},
		{[]int{1}, []int{1}, 1},
	}
}
