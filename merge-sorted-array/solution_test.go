package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	expected []int
	nums1    []int
	m        int
	nums2    []int
	n        int
}

func TestMerge(t *testing.T) {

	for _, dataSet := range getMergeData() {
		merge(dataSet.nums1, dataSet.m, dataSet.nums2, dataSet.n)
		assert.Equal(t, dataSet.expected, dataSet.nums1)
	}
}

func getMergeData() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[]int{1, 2, 2, 3, 5, 6}, []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{1}, []int{1}, 1, []int{}, 0},
		{[]int{1}, []int{0}, 0, []int{1}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3},
	}
}
