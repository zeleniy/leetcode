package unique_binary_search_trees_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	n      int
}

func TestGenerateTrees(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, len(generateTrees(dataSet.n)))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{1, 1},
		{2, 2},
		{5, 3},
		{14, 4},
		{42, 5},
	}
}
