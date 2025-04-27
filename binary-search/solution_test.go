package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SearchTestCase struct {
	expected, target int
	sequence         []int
}

func TestSearch(t *testing.T) {

	for _, dataSet := range getSearchData() {
		assert.Equal(t, dataSet.expected, Search(dataSet.sequence, dataSet.target))
	}
}

func getSearchData() []SearchTestCase {

	return []SearchTestCase{
		{-1, 3, []int{}},
		{0, 2, []int{2}},
		{-1, 3, []int{2}},
		{2, 3, []int{1, 2, 3}},
		{4, 9, []int{-1, 0, 3, 5, 9, 12}},
		{-1, 2, []int{-1, 0, 3, 5, 9, 12}},
		{-1, 0, []int{2, 5}},
	}
}
