package k_th_symbol_in_grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer int
	n      int
	k      int
}

func TestKthGrammar(t *testing.T) {

	for _, dataSet := range getTestDataSet() {
		assert.Equal(t, dataSet.answer, kthGrammar(dataSet.n, dataSet.k))
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{0, 1, 1},
		{0, 2, 1},
		{1, 2, 2},
		{0, 3, 1},
		{1, 3, 2},
		{1, 3, 3},
		{0, 3, 4},
	}
}
